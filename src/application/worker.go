package application

import (
	"fmt"
	"main/domain/entities"
	"main/domain/interfaces"
)

type WorkerApp struct {
	Worker        *entities.Worker
	WorkerService interfaces.WorkerService
	TaskService   interfaces.TaskService
}

func (app *WorkerApp) GetTasks() map[entities.TaskIdentifier]*entities.Task {
	return app.WorkerService.GetTasks(app.Worker)
}

func (app *WorkerApp) StartTask() {
	t := app.TaskService.CreateTask("TestTask")
	fmt.Printf("task: %v\n", t)

	app.WorkerService.AddTask(t, app.Worker)
}

func (app *WorkerApp) StopTask() {

}
