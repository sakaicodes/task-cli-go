package cmd

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/sakaicodes/task-cli-go/models"
)

var markOptions = []string{"To-Do", "In-Progress", "Done"}

/*
Marks status of a task in the task tracker. It accepts an ID and a new status (To-Do, In-Progress, Done). It loads existing tasks, finds the task by ID, updates its status, and saves it back to the file. If the ID or status is not provided, or if the status is invalid, it prints an error message and exits.
*/
func MarkStatus(args []string) {
	fs := flag.NewFlagSet("mark", flag.ExitOnError)
	id := fs.Int("id", 0, "ID of the task to update")
	status := fs.String("status", "", "New status for the task (To-Do, In-Progress, Done)")
	fs.Parse(args)

	*status = strings.Title(*status)

	// Check if ID is provided
	if *id == 0 {
		fmt.Println("error: --id is required")
		os.Exit(1)
	}

	// Check if status is provided
	if *status == "" {
		fmt.Println("error: --status is required")
		os.Exit(1)
	}

	//Check if status is valid
	if !slices.Contains(markOptions, *status) {
		fmt.Println("error: invalid status. Valid options are: To-Do, In-Progress, Done")
		os.Exit(1)
	}

	// Loads existing tasks
	tasks, err := models.LoadTasks()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	for i, task := range tasks {
		if task.ID == *id {
			tasks[i].Status = *status
			tasks[i].LastUpdated = time.Now().Format(time.RFC3339)
			err = models.SaveTasks(tasks)
			if err != nil {
				fmt.Println("error:", err)
				os.Exit(1)
			}
			fmt.Println("Task with ID", *id, "marked as", *status, "successfully.")
			return
		}
	}
	fmt.Println("error: task with ID", *id, "not found")
	os.Exit(1)
}
