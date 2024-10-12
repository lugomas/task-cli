# Task-CLI

`task-cli` is a simple command-line interface (CLI) tool for managing tasks stored in a JSON file. 
This tool allows you to create, update, delete, and manage the status of tasks efficiently, directly from your terminal.

## Features
- Add new tasks
- Update task descriptions
- Delete tasks
- Mark tasks as "in-progress", "done", or "todo"
- List all tasks or filter tasks by status

## Installation
To install task-cli, clone the repository and build the Go binary:

```bash
git clone https://github.com/lugomas/task-cli.git
cd task-cli
go build -o task-cli
```
You can now use the `task-cli` binary to manage your tasks.

Note: The Task ID is dynamically generated based on the current time, with a precision of seconds.

## Usage
```
### Adding a new task
./task-cli add "Buy groceries"
### Output: Task added successfully (ID: 2024101121521031)

### Updating and deleting tasks
./task-cli update 2024101121521031 "Buy groceries and cook dinner"
./task-cli delete 2024101121521031

### Marking a task as in progress or done
./task-cli mark-in-progress 2024101121521031
./task-cli mark-done 2024101121521031
./task-cli mark-todo 2024101121521031

### Listing all tasks
./task-cli list

### Listing tasks by status
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```

## License
This project is licensed under the MIT License.

## Project Inspiration
This project was developed based on the guidelines provided by [roadmap.sh's Task Tracker project](https://roadmap.sh/projects/task-tracker). 
The aim is to create a simple CLI tool for managing tasks, following best practices for command-line interfaces.