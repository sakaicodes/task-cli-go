package cmd

import (
	"fmt"
	"os"

	"github.com/sakaicodes/tasktracker/models"
)

/*
List all tasks in the task tracker. It loads existing tasks from the file and displays them in a readable format. If there is an error loading the tasks, it prints an error message and exits.
*/
func ListTasks() {
	// Load existing tasks
	tasks, err := models.LoadTasks()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	models.DisplayTasks(tasks)
}
