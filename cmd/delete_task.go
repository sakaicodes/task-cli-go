package cmd

import (
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/sakaicodes/tasktracker/models"
)

/*
Deletes a task from the task tracker. It takes an ID as a command-line argument, checks if the ID is provided, and then loads existing tasks from the file. It uses the slices.DeleteFunc function to create a new slice of tasks that excludes the task with the specified ID. If the length of the updated task list is the same as the original task list, it means that the task with the specified ID was not found, and an error message is printed. If the task is successfully deleted, the updated task list is saved back to the file, and a success message is printed.
*/
func DeleteTask(args []string) {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	id := fs.Int("id", 0, "ID of the task to delete")
	fs.Parse(args)

	// Check if ID is provided
	if *id == 0 {
		fmt.Println("error: --id is required")
		os.Exit(1)
	}

	// Loads existing tasks
	tasks, err := models.LoadTasks()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	// Deletes the task with the specified ID
	updatedTasks := slices.DeleteFunc(tasks, func(t models.Task) bool {
		return t.ID == *id
	})

	// If the length of updatedTasks is the same as tasks, it means the task with the specified ID was not found
	if len(updatedTasks) == len(tasks) {
		fmt.Println("error: task with ID", *id, "not found")
		os.Exit(1)
	}

	// Saves the updated task list back to the file
	err = models.SaveTasks(updatedTasks)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("Task with ID", *id, "deleted successfully.")

}
