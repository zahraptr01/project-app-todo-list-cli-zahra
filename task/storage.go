package task

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var filePath = filepath.Join("data", "tasks.json")

// function to load the task list
func LoadTasks() ([]Task, error) {
	var tasks []Task
	data, err := os.ReadFile(filePath)
	// Error Handling
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // empty the first time
		}
		return nil, err
	}
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

// Function to save task data into the tasks.json file.
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	// Error Handling
	if err != nil {
		return err
	}
	os.MkdirAll("data", os.ModePerm)
	return os.WriteFile(filePath, data, 0644)
}
