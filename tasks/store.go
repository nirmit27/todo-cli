/* Utility methods for READ and WRITE operations */
package tasks

import (
	"encoding/json"
	"os"
)

// JSON storage path
const TasksFile = "tasks.json"

func ReadTasks() ([]Task, error) {
	// Checking the JSON storage file's presence
	if _, err := os.Stat(TasksFile); os.IsNotExist(err) {
		if err := os.WriteFile(TasksFile, []byte("[]"), 0644); err != nil {
			return nil, err
		}
	}

	// Reading the JSON storage data
	data, err := os.ReadFile(TasksFile)
	if err != nil {
		return nil, err
	}

	// Conversion : JSON -> []Task
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func WriteTasks(tasks []Task) error {
	// Conversion : []Task -> JSON
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	// Write to the external JSON storage
	return os.WriteFile(TasksFile, data, 0644)
}
