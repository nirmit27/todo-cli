/* Utility functions for more specific operations */
package tasks

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

/* Misc. Operations */

// Validating input task status
func ValidStatus(status string) bool {
	if status == StatusDone || status == StatusInProgress || status == StatusTodo || status == "" {
		return true
	}
	return false
}

// Generating NEW 'Id' value
func GenerateID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].Id + 1
}

// Fetching 'task' by 'Id'
func FetchById(tasks []Task, id int) (*Task, error) {
	for i, task := range tasks {
		if task.Id == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("Task not found")
}

/* CRUD Operations */

// Adding a new 'task'
func AddTask(description string) error {
	if strings.TrimSpace(description) == "" {
		return fmt.Errorf("task description cannot be empty")
	}

	tasks, err := ReadTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		Id:          GenerateID(tasks),
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	err = WriteTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("\nTask added successfully! ID : %d\n\n", newTask.Id)
	return nil
}

// List all 'tasks'
func ListTasks(statusFilter string) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("\nNo tasks to show.")
		return nil
	}

	if !ValidStatus(statusFilter) {
		return fmt.Errorf("invalid status filter : %s", statusFilter)
	}

	counter := 0
	fmt.Println()

	for _, task := range tasks {
		if statusFilter == "" || statusFilter == task.Status {
			counter += 1
			fmt.Printf("%d. %s [%s]\t(Created: %s)\n", task.Id, task.Description, task.Status, task.CreatedAt.Format(TimeFormat))
		}
	}

	if counter == 0 {
		fmt.Printf("No tasks marked as : %s\n\n", statusFilter)
	} else if statusFilter != "" {
		fmt.Printf("\n%d task(s) marked as : %s\n\n", counter, statusFilter)
	} else {
		fmt.Printf("\nYou have %d task(s).\n\n", counter)
	}
	return nil
}

// Update a 'task'
func UpdateTask(description string, id int) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err
	}

	flag := true
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			flag = false
		}
	}

	if flag {
		return fmt.Errorf("task with ID : %d not found", id)
	}

	err = WriteTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("\nTask [%d] updated successfully!\n\n", id)
	return nil
}

// Update 'task' status / MARK a 'task'
func MarkTask(id int, status string) error {
	if !ValidStatus(status) {
		return fmt.Errorf("invalid status filter : %s", status)
	}

	tasks, err := ReadTasks()
	if err != nil {
		return err
	}

	flag := true
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()

			flag = false
			break
		}
	}

	if flag {
		return fmt.Errorf("task with ID : %d not found", id)
	}

	err = WriteTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("\nTask [%d] MARKED as %s.\n\n", id, status)
	return nil
}

// Delete a 'task'
func DeleteTask(id int) error {
	tasks, err := ReadTasks()
	if err != nil {
		return err
	}

	index := -1
	for i, task := range tasks {
		if task.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("task with ID : %d not found", id)
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	err = WriteTasks(tasks)
	if err != nil {
		return err
	}

	fmt.Printf("\nTask [%d] deleted successfully!\n\n", id)
	return nil
}
