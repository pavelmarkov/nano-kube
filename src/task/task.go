package task

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"main/dockerclient"
	"time"

	"github.com/google/uuid"
)

type TaskState int

const (
	Pending TaskState = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	TaskId uuid.UUID
	Name   string
	State  TaskState
}

func CreateTask(name string) *Task {
	t := &Task{
		TaskId: uuid.New(),
		State:  Pending,
		Name:   name,
	}

	return t
}

func RunTask(t *Task) error {

	DockerClient := &dockerclient.DockerClient{}

	ContainerConfig := dockerclient.ContainerConfig{
		ImageName:     "alpine",
		ContainerName: t.Name,
		Commands:      []string{"sh"},
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()

	cli, err := dockerclient.GetDockerClient(DockerClient, &ctx, &ContainerConfig)
	if err != nil {
		log.Fatalf("Failed to get docker client: %v", err)
		return err
	}
	defer cli.Close()

	err = dockerclient.PullDockerImage(DockerClient)
	if err != nil {
		slog.Error("Failed to pull image", "error", err)
		return err
	}

	_, err = dockerclient.CreateDockerContainer(DockerClient)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
		return err
	}

	err = dockerclient.StartDockerContainer(DockerClient)
	if err != nil {
		log.Fatalf("Failed to start container: %v", err)
		return err
	}

	err = dockerclient.AttachDockerStdout(DockerClient)
	if err != nil {
		log.Fatalf("Failed to attach: %v", err)
		return err
	}

	return nil
}

// Stringer interfaces

func (t Task) String() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.State)
}
func (s TaskState) String() string {
	return []string{"Pending", "Scheduled", "Running", "Completed", "Failed"}[s]
}
