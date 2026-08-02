package dockerclient

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type ContainerConfig struct {
	ImageName     string
	ContainerName string
	Commands      []string
}

type DockerClient struct {
	ctx    context.Context
	cli    *client.Client
	config *ContainerConfig
	resp   *container.CreateResponse
}

func GetDockerClient(
	dc *DockerClient,
	ctx *context.Context,
	config *ContainerConfig,
) (*client.Client, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("Failed to get current user: %v", err)
	}

	socketPath := "unix://" + usr.HomeDir + "/.docker/run/docker.sock"
	log.Printf("Attempting to connect to Docker at: %s", socketPath)

	cli, err := client.NewClientWithOpts(
		client.WithHost(socketPath),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Docker client: %v", err)
	}
	defer cli.Close()

	_, err = cli.Ping(*ctx)
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to Docker daemon: %v\nMake sure Docker Desktop is running!", err)
	}

	log.Println("✅ Connected to Docker daemon")

	dc.ctx = *ctx
	dc.cli = cli
	dc.config = config

	return cli, nil
}

func PullDockerImage(dc *DockerClient) (err error) {
	log.Printf("📥 Pulling image: %s...\n", dc.config.ImageName)
	reader, err := dc.cli.ImagePull(dc.ctx, dc.config.ImageName, image.PullOptions{})
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)
	reader.Close()
	return nil
}

func CreateDockerContainer(dc *DockerClient) (*container.CreateResponse, error) {
	resp, err := dc.cli.ContainerCreate(
		dc.ctx,
		&container.Config{
			Image:        dc.config.ImageName,
			Cmd:          dc.config.Commands,
			Tty:          true,
			AttachStdout: true,
			AttachStderr: true,
			AttachStdin:  true,
			OpenStdin:    true,
		},
		&container.HostConfig{
			AutoRemove: true,
		},
		nil, nil, dc.config.ContainerName,
	)
	if err != nil {
		return nil, err
	}

	log.Printf("Container created: %s", resp.ID[:12])

	dc.resp = &resp

	return &resp, nil
}

func StartDockerContainer(dc *DockerClient) (err error) {
	err = dc.cli.ContainerStart(dc.ctx, dc.resp.ID, container.StartOptions{})
	if err != nil {
		return err
	}
	return nil
}

func AttachDockerStdout(dc *DockerClient) (err error) {

	// Attach to the container's streams for interactive session
	attachResp, err := dc.cli.ContainerAttach(dc.ctx, dc.resp.ID, container.AttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: true,
		Stderr: true,
	})
	if err != nil {
		return err
	}
	defer attachResp.Close()

	fmt.Printf("✅All is ready! Type 'exit' to quit.\n")

	go io.Copy(attachResp.Conn, os.Stdin)
	io.Copy(os.Stdout, attachResp.Conn)

	return nil

}
