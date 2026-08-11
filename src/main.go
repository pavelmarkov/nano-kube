package main

import (
	"fmt"
	"main/api"
	"main/worker"
)

func main() {
	w := worker.CreateWorker("TestWorker")
	fmt.Printf("worker: %v\n", w)

	go w.RunTasks()

	a := api.Api{Worker: w}
	a.Init()
	a.Router.Run()
}
