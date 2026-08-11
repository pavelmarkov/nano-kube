package worker

import (
	"fmt"
	"log"
	"main/task"
	"time"

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
	for {
		t := w.TaskQueue.Dequeue()
		if t == nil {
			log.Println("Queue is empty")
		} else {
			firstTaskInQueue := t.(*task.Task)
			err := task.RunTask(firstTaskInQueue)
			if err != nil {
				log.Printf("Failed running task: %v", err)
				w.TaskStorage[firstTaskInQueue.TaskId].State = task.Failed
			} else {
				w.TaskStorage[firstTaskInQueue.TaskId].State = task.Running
			}
		}
		log.Println("Sleep for 10 seconds")
		time.Sleep(10 * time.Second)
	}
}

func (w *Worker) GetTasks() map[task.TaskIdentifier]*task.Task {
	return w.TaskStorage
}

// Stringer interfaces
func (w *Worker) String() string {
	return fmt.Sprintf("%s (%d)", w.Name, len(w.TaskStorage))
}
