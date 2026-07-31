package main

import (
	"fmt"
	"main/task"
)

func main() {
	t := task.CreateTask("Test Task")
	fmt.Printf("task: %v\n", t)
}
