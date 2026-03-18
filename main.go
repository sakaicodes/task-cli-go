package main

import (
	"fmt"
	"os"

	"github.com/sakaicodes/tasktracker/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("error: command is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		cmd.AddTasks(os.Args[2:])
	case "list":
		cmd.ListTasks()
	case "update":
		cmd.UpdateTasks(os.Args[2:])
	case "delete":
		cmd.DeleteTasks(os.Args[2:])
	default:
		fmt.Println("error: unknown command")
		os.Exit(1)
	}
}
