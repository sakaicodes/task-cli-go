package cmd

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/sakaicodes/tasktracker/models"
)

var markOptions = []string{"To-Do", "In-Progress", "Done"}

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
