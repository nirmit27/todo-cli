package cmds

import (
	"fmt"
)

// Printing the help menu
func PrintHelp() {
	fmt.Print("\n\t⚒️   Task Tracking CLI  📝\n\n")
	fmt.Println("Usage: task-cli <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  add <description>          - Add new task")
	fmt.Println("  list [status]              - List tasks (optionally filtered by status)")
	fmt.Println("  update <id> <description>  - Update task description by ID")
	fmt.Println("  delete <id>                - Delete task by ID")
	fmt.Println("  mark <id> <status>         - Mark task status (todo, in-progress, done)")
}
