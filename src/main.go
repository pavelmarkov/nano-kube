package main

import (
	"fmt"
	"main/task"
	"main/worker"
)

func main() {
	w := worker.CreateWorker("TestWorker")
	fmt.Printf("worker: %v\n", w)

	t := task.CreateTask("TestTask")
	fmt.Printf("task: %v\n", t)

	w.AddTask(t)

	w.RunTasks()
}
