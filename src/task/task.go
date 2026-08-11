package task

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"main/dockerclient"
	"time"

	"encoding/json"

	"github.com/google/uuid"
)

type TaskState int
type TaskIdentifier = uuid.UUID

const (
	Pending TaskState = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	TaskId TaskIdentifier
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

// Stringer interfaces
func (t *Task) String() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.State)
}
func (s TaskState) String() string {
	return []string{"Pending", "Scheduled", "Running", "Completed", "Failed"}[s]
}

func (s TaskState) MarshalJSON() ([]byte, error) {
	var taskState string
	switch s {
	case Pending:
		taskState = "Pending"
	case Scheduled:
		taskState = "Scheduled"
	case Running:
		taskState = "Running"
	case Completed:
		taskState = "Completed"
	case Failed:
		taskState = "Failed"
	default:
		taskState = "Unknown"
	}
	return json.Marshal(taskState)
}
