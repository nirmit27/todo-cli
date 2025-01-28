package cmds

import (
	"fmt"
	"strings"
	"todo-cli/tasks"
)

func Add(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage : todo-cli add <description>\n - Add a new task with a description.")
		return
	}

	description := strings.Join(args, " ")
	if strings.TrimSpace(description) == "" {
		fmt.Print("\nTask description cannot be empty. Please try again.\n\n")
		return
	}

	if err := tasks.AddTask(description); err != nil {
		fmt.Println("Error adding task :", err)
	}
}
