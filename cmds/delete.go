package cmds

import (
	"fmt"
	"strconv"
	"todo-cli/tasks"
)

func Delete(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage : todo-cli delete <id>\n - Delete a task by ID.")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Print("\nTask ID must be a positive number. Please try again.\n\n")
		return
	}

	if err := tasks.DeleteTask(id); err != nil {
		fmt.Printf("\nError deleting task : %s\n\n", err)
	}
}
