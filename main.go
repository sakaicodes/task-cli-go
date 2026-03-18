package main

import (
	"fmt"
	"os"

	"github.com/sakaicodes/task-cli-go/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: command is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		cmd.AddTask(os.Args[2:])
	case "list":
		cmd.ListTasks()
	case "update":
		cmd.UpdateTask(os.Args[2:])
	case "mark":
		cmd.MarkStatus(os.Args[2:])
	case "delete":
		cmd.DeleteTask(os.Args[2:])
	default:
		fmt.Println("error: unknown command")
		os.Exit(1)
	}
}
