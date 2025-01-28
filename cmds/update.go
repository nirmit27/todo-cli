package cmds

import (
	"fmt"
	"strconv"
	"strings"
	"todo-cli/tasks"
)

func Update(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage : todo-cli update <int> <description>\n - Update task description by ID.")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Print("\nTask ID must be a positive number. Please try again.\n\n")
		return
	}

	description := strings.Join(args[1:], " ")
	if strings.TrimSpace(description) == "" {
		fmt.Print("\nNew description cannot be empty. Please try again.\n\n")
		return
	}

	if err := tasks.UpdateTask(description, id); err != nil {
		fmt.Printf("\nError updating task : %s\n\n", err)
	}
}
