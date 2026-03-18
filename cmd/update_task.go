package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/sakaicodes/task-cli-go/models"
)

/*
Updates the title of an existing task in the task tracker. It takes an ID and a new title as command-line arguments, checks if both are provided, and then loads existing tasks from the file. It iterates through the tasks to find the one with the specified ID, updates its title, and saves the updated task list back to the file. If the task with the specified ID is not found, it prints an error message.
*/
func UpdateTask(args []string) {
	fs := flag.NewFlagSet("update", flag.ExitOnError)
	id := fs.Int("id", 0, "ID of the task to update")
	newDescription := fs.String("title", "", "New title for the task")
	fs.Parse(args)

	//Checks if new title is provided
	if *newDescription == "" {
		fmt.Println("error: --title is required")
		os.Exit(1)
	}

	//Checks if ID is provided
	if *id == 0 {
		fmt.Println("error: --id is required")
		os.Exit(1)

	}

	// Load existing tasks
	tasks, err := models.LoadTasks()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	// Find the task by ID and update its title
	for i, task := range tasks {
		if task.ID == *id {
			tasks[i].Title = *newDescription
			tasks[i].LastUpdated = time.Now().Format(time.RFC3339)
			err = models.SaveTasks(tasks)
			if err != nil {
				fmt.Println("error:", err)
				os.Exit(1)
			}
			fmt.Println("Task with ID", *id, "updated successfully.")
			return
		}
	}
	fmt.Println("error: task with ID", *id, "not found")
	os.Exit(1)
}
