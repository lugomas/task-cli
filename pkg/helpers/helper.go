package helpers

import "fmt"

func ShowHelp() {
	helpText := `Usage: task-cli <command>

Commands:
  helper                      Show this helper message
  list                      Lists all task
  list done                 Lists task with status done
  list in-progress          Lists task with status in-progress
  list todo                 Lists task with status todo
  add <task description>     Add task with status todo
  delete <taskID>            Delete task
  update <taskID> <task description>  Update task description
  mark-done <taskID>         Mark task with status done
  mark-in-progress <taskID>  Mark task with status in-progress
  mark-todo <taskID>         Mark task with status todo`

	fmt.Println(helpText)
}

var unknownCommand = "Unknown command"
var taskHelp = "Use 'task-tracker help' to see available commands."

func UnknownCommand() {
	fmt.Println(unknownCommand)
	fmt.Println(taskHelp)
}
