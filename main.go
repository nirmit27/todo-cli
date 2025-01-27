package main

import (
	"fmt"
	"os"
	"todo-cli/cmds"
)

func main() {
	// Display the menu
	if len(os.Args) < 2 {
		cmds.PrintHelp()
		return
	}

	// Parsing the commands
	command := os.Args[1]
	args := os.Args[2:]

	// Command handlers
	switch command {
	case "add":
		cmds.Add(args)
	case "list":
		cmds.List(args)
	case "update":
		cmds.Update(args)
	case "delete":
		cmds.Delete(args)
	case "mark":
		cmds.Mark(args)
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		cmds.PrintHelp()
	}
}
