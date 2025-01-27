package cmds

import (
	"fmt"
	"todo-cli/tasks"
)

func List(args []string) {
	status := ""
	if len(args) > 0 {
		status = args[0]
	}

	if err := tasks.ListTasks(status); err != nil {
		fmt.Println("Error:", err)
	}
}
