package task

import (
	"fmt"

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

// Stringer interfaces

func (t Task) String() string {
	return fmt.Sprintf("%s (%s)", t.Name, t.State)
}
func (s TaskState) String() string {
	return []string{"Pending", "Scheduled", "Running", "Completed", "Failed"}[s]
}
