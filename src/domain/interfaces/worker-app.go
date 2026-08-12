package interfaces

import "main/domain/entities"

type WorkerApp interface {
	GetTasks() map[entities.TaskIdentifier]*entities.Task
	StartTask()
	StopTask()
}
