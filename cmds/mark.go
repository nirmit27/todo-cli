package cmds

import (
	"fmt"
	"strconv"
	"strings"
	"todo-cli/tasks"
)

func Mark(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage : todo-cli mark <id> <status>\n - Mark task status (todo, in-progress, done)")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Print("\nTask ID must be a positive number. Please try again.\n\n")
		return
	}

	status := strings.TrimSpace(args[1])
	if err := tasks.MarkTask(id, status); err != nil {
		fmt.Printf("\nError fetching task : %s\n\n", err)
	}
}
