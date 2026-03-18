package cmd

import (
	"fmt"
	"os"

	"github.com/sakaicodes/tasktracker/models"
)

func ListTasks() {
	// Load existing tasks
	tasks, err := models.LoadTasks()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	models.DisplayTasks(tasks)
}
