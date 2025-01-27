package cmds

import (
	"fmt"
)

func Delete(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage : todo-cli delete <id>\n - Delete a task by ID.")
		return
	}

	// TODO
	
}
