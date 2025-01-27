package cmds

import (
	"fmt"
)

func Mark(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage : todo-cli mark <id> <status>\n - Mark task status (todo, in-progress, done)")
		return
	}

	// TODO

}
