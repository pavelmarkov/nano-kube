package main

import (
	"fmt"
	"main/task"
)

func main() {
	t := task.CreateTask("TestTask")
	fmt.Printf("task: %v\n", t)

	task.RunTask(t)
}
