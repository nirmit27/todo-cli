package cmds

import (
	"fmt"
)

func PrintHelp() {
	fmt.Println(`
⚒️   Task Tracking CLI 📝

Usage: task-cli <command> [arguments]

Commands:
  add <description>                 - Add new task
  list [status]                     - List tasks (optionally filtered by status)
  update <id> <description>         - Update task description by ID
  delete <id>                       - Delete task by ID
  mark <id> <status>                - Mark task status (todo, in-progress, done)
	`)
}
