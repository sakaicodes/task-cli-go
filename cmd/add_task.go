package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/sakaicodes/task-cli-go/models"
)

/*
Add tasks to the task tracker. It accepts a title and an optional status (default is ""). It loads existing tasks, creates a new task, and saves it back to the file. If the title is not provided, it prints an error message and exits.
*/
func AddTask(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	title := fs.String("title", "", "Title of the task")
	status := fs.String("status", "Unassigned", "Status of the task")
	fs.Parse(args)

	// Checks if title is provided
	if *title == "" {
		fmt.Println("error: --title is required")
		os.Exit(1)
	}

	// Load existing tasks
	tasks, err := models.LoadTasks()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	// Creates a new task and saves it
	newTask := models.CreateTask(tasks, *title, *status)
	tasks = append(tasks, *newTask)
	err = models.SaveTasks(tasks)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
