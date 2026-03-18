package models

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status int    `json:"status"`
}

const filePath = "tasks.json"

func CreateTask(tasks []Task, title string, status int) *Task {
	task := Task{
		ID:     nextID(tasks),
		Title:  title,
		Status: status,
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
		data = append(data, []string{strconv.Itoa(task.ID), task.Title, strconv.Itoa(task.Status)})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Title", "Status"})
	table.Bulk(data)
	table.Render()
}
