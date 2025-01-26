package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Constants
const (
	timeFormat = time.RFC1123
	StatusTodo = "todo"
	StatusInProgress = "in-progress"
	StatusDone = "done"
)

// Struct for 'Task'
type Task struct {
	Id int
	Description string
	Status string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Temporary storage for tasks
var tasks []Task
var nextId int = 1

// Validating input task status
func validStatus(status string) bool {
	if status == StatusDone || status == StatusInProgress || status == StatusTodo {
		return true
	}
	return false
}

// Sanitizing user input strings
func readInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("Failed to read input.")
	}
	return strings.TrimSpace(input), nil
}

// Sanitizing user input for task Id
func readId(reader *bufio.Reader) (int, error) {
	idStr, error := readInput("\nEnter task ID : ", reader)
	if error != nil {
		return 0, error
	}

	id, error := strconv.Atoi(idStr)
	if error != nil {
		return 0, errors.New("Invalid input. ID should be a positive number.")
	}
	return id, nil
}

/* CRUD Operations */

// Adding a task
func addTask(reader *bufio.Reader) {
	fmt.Print("\n----- Add a Task -----\n")

	desc, err := readInput("\nEnter description : ", reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	if desc == "" {
		fmt.Println("Task description cannot be empty!")
		return
	}

	task := Task{
		Id: nextId,
		Description: desc,
		Status: StatusTodo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{}, // default value
	}

	tasks = append(tasks, task)
	nextId++
	fmt.Println("\nTask added successfully! 👍")
}

// List all tasks
func listAllTasks() {
	if len(tasks) == 0 {
		fmt.Println("\nNo tasks to show.")
		return
	}

	fmt.Print("\n----- All Tasks -----\n\n")
	for _, task := range tasks {
		fmt.Printf("%d. %s [%s]\t< %v >\n", task.Id, task.Description, task.Status, task.CreatedAt.Format(timeFormat))
	}
}

// Listing tasks by STATUS
func listTasksByStatus(reader *bufio.Reader) {
	if len(tasks) == 0 {
		fmt.Println("\nNo tasks to show.")
		return
	}
	fmt.Print("\n--- Tasks by Status ---\n")
	
	status, error := readInput("\nEnter task status : ", reader)
	if error != nil {
		fmt.Println(error)
		return
	}

	if !validStatus(status) {
		fmt.Println("\nInvalid input. Please try again!")
		return
	}

	flag := true
	fmt.Printf("\nStatus : [%s]\n\n", status)
	for _, task := range tasks{
		if task.Status == status {
			flag = false
			fmt.Printf("%d. %s [%s]\t< %v >\n", task.Id, task.Description, task.Status, task.CreatedAt.Format(timeFormat))
		}
	}
	if (flag) {
		fmt.Printf("No tasks have been marked as %s.\n", status)
	}
}

// Update task
func updateTask(reader *bufio.Reader) {
	id, error := readId(reader)
	if error != nil {
		fmt.Println(error)
		return
	}

	updatedDesc, error := readInput("\nEnter NEW description : ", reader)
	if error != nil {
		fmt.Println(error)
		return
	}

	flag := true
	for i, task := range tasks {
		if task.Id == id {
			flag = false
			tasks[i].Description = updatedDesc

			fmt.Println("\nTask UPDATED successfully 📝")
			return
		}
	}
	if flag {
		fmt.Println("No tasks found with ID :", id)
	}
}

// Delete task
func deleteTask(reader *bufio.Reader) {
	id, error := readId(reader)
	if error != nil {
		fmt.Println(error)
		return
	}

	flag := true
	for i, task := range tasks {
		if task.Id == id {
			flag = false
			tasks = append(tasks[:i], tasks[i+1:]...)

			fmt.Println("\nTask DELETED successfully 🗑️")
			return
		}
	}
	if flag {
		fmt.Println("No tasks found with ID :", id)
	}
}

// Mark as DONE
func markAsDone(reader *bufio.Reader) {
	id, error := readId(reader)
	if error != nil {
		fmt.Println(error)
		return
	}

	flag := true
	for i, task := range tasks {
		if task.Id == id {
			if task.Status == StatusDone {
				fmt.Println("\nThis task has already been DONE.")
			} else {
				tasks[i].Status = StatusDone
				fmt.Println("\nTask marked as DONE ☑️")
			}
			return
		}
	}
	if flag {
		fmt.Println("No tasks found with ID :", id)
	}
}

// Mark as IN-PROGRESS
func markAsInProgress(reader *bufio.Reader) {
	id, error := readId(reader)
	if error != nil {
		fmt.Println(error)
		return
	}

	flag := true
	for i, task := range tasks {
		if task.Id == id {
			if task.Status == StatusInProgress {
				fmt.Println("\nThis task is already IN-PROGRESS.")
			} else {
				tasks[i].Status = StatusInProgress
				fmt.Println("\nTask marked as IN-PROGRESS 🧑‍🏭")
			}
			return
		}
	}
	if flag {
		fmt.Println("No tasks found with ID :", id)
	}
}


/* Driver method */
func main(){
	// TODO: Panic handler
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nError: %v\n", r)
			fmt.Println("Please try again!")
		}
	}()

	fmt.Print("\n ⚒️   Task Tracking CLI  📝    ")

	for {
		// Main Menu
		fmt.Println("\n\n-------- MAIN MENU --------")
		fmt.Println("\n [1] Add a task")
		fmt.Println(" [2] List all tasks")
		fmt.Println(" [3] List tasks by STATUS")
		fmt.Println(" [4] Update a task")
		fmt.Println(" [5] Delete a task")
		fmt.Println(" [6] Mark as IN-PROGRESS")
		fmt.Println(" [7] Mark as DONE")
		fmt.Println(" [8] Exit")
		fmt.Println("\n---------------------------")

		// Input logic
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nEnter your option: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		// Cleaning the input
		input = strings.TrimSpace(input)
		option, err := strconv.Atoi(input)

		if err != nil || option < 1 || option > 8 {
			fmt.Println("\nInvalid option! Please enter a number between 1 and 8.")
			continue
		}

		// Handlers for options
		switch option {
		case 1:
			addTask(reader)
		case 2:
			listAllTasks()
		case 3:
			listTasksByStatus(reader)
		case 4:
			updateTask(reader)
		case 5:
			deleteTask(reader)
		case 6:
			markAsInProgress(reader)
		case 7:
			markAsDone(reader)
		case 8:
			fmt.Print("\nGoodbye! 👋\n\n")
			return
		}
	}
}
