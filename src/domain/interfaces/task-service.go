package interfaces

import "main/domain/entities"

type TaskService interface {
	CreateTask(name string) *entities.Task
	RunTask(t *entities.Task) error
}
