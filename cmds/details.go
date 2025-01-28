package cmds

import (
	"fmt"
	"strconv"
	"todo-cli/tasks"
)

func Details(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage : todo-cli task <id>\n - View a task by ID")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Print("\nTask ID must be a positive number. Please try again.\n\n")
		return
	}

	if err = tasks.FetchDetails(id); err != nil {
		fmt.Printf("\nError : %s\n\n", err)
	}
}
