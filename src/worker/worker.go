package worker

import (
	"fmt"
	"log"
	"main/task"

	"github.com/golang-collections/collections/queue"
)

type Worker struct {
	Name        string
	TaskQueue   queue.Queue
	TaskStorage map[task.TaskIdentifier]*task.Task
}

func CreateWorker(name string) *Worker {
	w := &Worker{
		Name:        name,
		TaskQueue:   *queue.New(),
		TaskStorage: make(map[task.TaskIdentifier]*task.Task),
	}

	return w
}

func (w *Worker) AddTask(t *task.Task) {
	w.TaskQueue.Enqueue(t)
	w.TaskStorage[t.TaskId] = t
}

func (w *Worker) RunTasks() {
	t := w.TaskQueue.Dequeue()
	if t == nil {
		log.Println("Queue is empty")
		return
	}
	firstTaskInQueue := t.(*task.Task)
	task.RunTask(firstTaskInQueue)
}

// Stringer interfaces
func (w *Worker) String() string {
	return fmt.Sprintf("%s (%d)", w.Name, len(w.TaskStorage))
}
