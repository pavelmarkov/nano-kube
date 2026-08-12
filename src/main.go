package main

import (
	"fmt"
	"main/application"
	"main/presentation/ginapp"
	"main/services"
)

func main() {
	ts := services.TaskService{}

	ws := services.WorkerService{
		TaskService: &ts,
	}

	w := ws.CreateWorker("TestWorker")
	fmt.Printf("worker: %v\n", w)
	go ws.RunTasks(w)

	wa := application.WorkerApp{
		Worker:        w,
		WorkerService: &ws,
		TaskService:   &ts,
	}

	a := ginapp.Api{
		WorkerApp: &wa,
	}
	a.Init()
	a.Router.Run()
}
