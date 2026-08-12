package interfaces

import (
	"main/domain/entities"
)

type WorkerService interface {
	CreateWorker(name string) *entities.Worker
	AddTask(t *entities.Task, w *entities.Worker)
	RunTasks(w *entities.Worker)
	GetTasks(w *entities.Worker) map[entities.TaskIdentifier]*entities.Task
}
