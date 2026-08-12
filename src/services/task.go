package services

import (
	"context"
	"log"
	"log/slog"
	"main/domain/entities"
	"main/infrastructure"
	"time"

	"github.com/google/uuid"
)

type TaskService struct{}

func (service *TaskService) CreateTask(name string) *entities.Task {
	t := &entities.Task{
		TaskId: uuid.New(),
		State:  entities.Pending,
		Name:   name,
	}

	return t
}

func (service *TaskService) RunTask(t *entities.Task) error {

	DockerClient := &infrastructure.DockerClient{}

	ContainerConfig := infrastructure.ContainerConfig{
		ImageName:     "strm/helloworld-http",
		ContainerName: t.Name,
		Commands:      []string{},
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelFunc()

	cli, err := DockerClient.GetDockerClient(&ctx, &ContainerConfig)
	if err != nil {
		log.Printf("Failed to get docker client: %v", err)
		return err
	}
	defer cli.Close()

	err = DockerClient.PullDockerImage()
	if err != nil {
		slog.Error("Failed to pull image", "error", err)
		return err
	}

	_, err = DockerClient.CreateDockerContainer()
	if err != nil {
		log.Printf("Failed to create container: %v", err)
		return err
	}

	err = DockerClient.StartDockerContainer()
	if err != nil {
		log.Printf("Failed to start container: %v", err)
		return err
	}

	// err = DockerClient.AttachDockerStdout()
	// if err != nil {
	// 	log.Printf("Failed to attach: %v", err)
	// 	return err
	// }

	return nil
}
