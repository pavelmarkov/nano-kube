package entities

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
)

type Worker struct {
	Name        string
	TaskQueue   queue.Queue
	TaskStorage map[TaskIdentifier]*Task
}

// Stringer interfaces
func (w *Worker) String() string {
	return fmt.Sprintf("%s (%d)", w.Name, len(w.TaskStorage))
}
