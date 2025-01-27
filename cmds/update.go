package cmds

import (
	"fmt"
	"strconv"
	"strings"
	"todo-cli/tasks"
)

func Update(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage : todo-cli update <int> <description>\n - Update task description by ID.")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("\nTask ID must be a positive number. Please try again.")
		return
	}

	description := strings.Join(args[1:], " ")

	if err := tasks.UpdateTask(description, id); err != nil {
		fmt.Println("\nError updating task :", err)
	}
}
