package cmd

import (
	"fmt"
	"roadmaps/projects/task-tracker/task-cli/pkg/helpers"
)

func HandleCommands(args []string) {

	switch args[1] {
	default:
		helpers.UnknownCommand()

	case "help":
		helpers.ShowHelp()
		return

	case "list":
		if len(args) == 2 {
			tasks, err := ReadAllTasks()
			if err != nil {
				fmt.Println(err)
				return
			}
			PrintTasks(tasks)
			return
		}
		if len(args) == 3 {
			tasks, err := ReadTasksStatus(args[2])
			if err != nil {
				helpers.UnknownCommand()
				return
			}
			if tasks == nil {
				fmt.Println("No tasks found")
				return
			}
			PrintTasks(tasks)
			return
		}
		helpers.UnknownCommand()
		return

	case "add":
		if len(args) == 3 {
			taskTitle := args[2]
			taskId, err := WriteTask(taskTitle)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Task added successfully (ID: %v)\n", taskId)
			return
		}
		helpers.UnknownCommand()
		return

	case "update":
		if len(args) == 4 {
			taskId := args[2]
			taskTitle := args[3]
			err := UpdateDescription(taskId, taskTitle)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Task updated successfully (ID: %v)\n", taskId)
			return
		}
		helpers.UnknownCommand()
		return

	case "delete":
		if len(args) == 3 {
			taskId := args[2]
			err := DeleteTask(taskId)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Task deleted successfully (ID: %v)\n", taskId)
			return
		}
		helpers.UnknownCommand()
		return

	case "mark-done":
		if len(args) == 3 {
			taskId := args[2]
			err := MarkDone(taskId)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Task updated successfully (ID: %v)\n", taskId)
			return
		}
		helpers.UnknownCommand()
		return

	case "mark-in-progress":
		if len(args) == 3 {
			taskId := args[2]
			err := MarkInProgress(taskId)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Task updated successfully (ID: %v)\n", taskId)
			return
		}
		helpers.UnknownCommand()
		return

	case "mark-todo":
		if len(args) == 3 {
			taskId := args[2]
			err := MarkToDo(taskId)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Task updated successfully (ID: %v)\n", taskId)
			return
		}
		helpers.UnknownCommand()
		return
	}
}
