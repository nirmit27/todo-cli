package cmds

import (
	"fmt"
)

func Update(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage : todo-cli update <int> <description>\n - Update task description by ID.")
		return
	}

	// TODO

}
