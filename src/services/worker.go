package services

import (
	"log"
	"main/domain/entities"
	"time"

	"github.com/golang-collections/collections/queue"
)

type WorkerService struct {
	TaskService *TaskService
}

func (service *WorkerService) CreateWorker(name string) *entities.Worker {
	w := &entities.Worker{
		Name:        name,
		TaskQueue:   *queue.New(),
		TaskStorage: make(map[entities.TaskIdentifier]*entities.Task),
	}

	return w
}

func (service *WorkerService) AddTask(t *entities.Task, w *entities.Worker) {
	w.TaskQueue.Enqueue(t)
	w.TaskStorage[t.TaskId] = t
}

func (service *WorkerService) RunTasks(w *entities.Worker) {
	for {
		t := w.TaskQueue.Dequeue()
		if t == nil {
			log.Println("Queue is empty")
		} else {
			firstTaskInQueue := t.(*entities.Task)
			err := service.TaskService.RunTask(firstTaskInQueue)
			if err != nil {
				log.Printf("Failed running task: %v", err)
				w.TaskStorage[firstTaskInQueue.TaskId].State = entities.Failed
			} else {
				w.TaskStorage[firstTaskInQueue.TaskId].State = entities.Running
			}
		}
		log.Println("Sleep for 10 seconds")
		time.Sleep(10 * time.Second)
	}
}

func (service *WorkerService) GetTasks(w *entities.Worker) map[entities.TaskIdentifier]*entities.Task {
	return w.TaskStorage
}
