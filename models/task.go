package models

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	LastUpdated string `json:"last_updated"`
}

const filePath = "tasks.json"

func CreateTask(tasks []Task, title string, status string) *Task {
	task := Task{
		ID:          nextID(tasks),
		Title:       title,
		Status:      status,
		CreatedAt:   time.Now().Format(time.RFC3339),
		LastUpdated: time.Now().Format(time.RFC3339),
	}
	return &task
}

func nextID(tasks []Task) int {
	max := 0
	for _, task := range tasks {
		if task.ID > max {
			max = task.ID
		}
	}
	return max + 1
}

func LoadTasks() ([]Task, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // No file yet, return empty slice
		}
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644) // 0644 is the file permission
}

func DisplayTasks(tasks []Task) {
	data := [][]string{}
	for _, task := range tasks {
		data = append(data, []string{strconv.Itoa(task.ID), task.Title, task.Status, task.CreatedAt, task.LastUpdated})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Title", "Status", "Created At", "Last Updated"})
	table.Bulk(data)
	table.Render()
}
