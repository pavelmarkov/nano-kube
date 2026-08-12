package entities

import (
	"encoding/json"
	"fmt"

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
