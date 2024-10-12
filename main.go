package main

import (
	"os"
	"roadmaps/projects/task-tracker/task-cli/cmd"
	"roadmaps/projects/task-tracker/task-cli/pkg/helpers"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		helpers.UnknownCommand()
		return
	}
	cmd.HandleCommands(args)
}
