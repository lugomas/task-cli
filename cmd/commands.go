package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id          string    `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

const taskFile = "json.txt"

func PrintTasks(tasks []Task) {
	for i, task := range tasks {
		fmt.Printf("%d. ID: %s\n", i+1, task.Id)
		fmt.Printf("   Description: %s\n", task.Description)
		fmt.Printf("   Status: %s\n", task.Status)
		fmt.Printf("   Created At: %s\n", task.CreatedAt.Format(time.RFC3339)) // Format time
		fmt.Printf("   Updated At: %s\n", task.UpdatedAt.Format(time.RFC3339)) // Format time
		fmt.Println("   --------------------")
	}
}

func checkFileExistence() error {
	_, err := os.Stat(taskFile)
	if os.IsNotExist(err) {
		return errors.New("No task were found\n Use 'task-cli add' to create a task.\n")
	}
	return nil
}

func readFile() ([]Task, error) {

	fileData, err := os.ReadFile(taskFile)
	if err != nil {
		return nil, fmt.Errorf("error reading task file: %w", err)
	}

	var tasks []Task
	err = json.Unmarshal(fileData, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling tasks: %w", err)
	}

	return tasks, nil
}

func ReadAllTasks() ([]Task, error) {

	err := checkFileExistence()
	if err != nil {
		return nil, err
	}

	return readFile()
}

func createFile(tasks []Task) error {
	newData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling tasks: %w", err)
	}
	err = os.WriteFile(taskFile, newData, 0644)
	if err != nil {
		return fmt.Errorf("error writing to task file: %w", err)
	}
	return nil
}

func WriteTask(taskTitle string) (string, error) {

	// Generate a unique taskID based on time
	taskID := time.Now().Format("200601021504105")
	newTask := Task{
		Id:          taskID,
		Description: taskTitle,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	var tasks []Task
	err := checkFileExistence()
	if err != nil {
		tasks = append(tasks, newTask)
		err := createFile(tasks)
		if err != nil {
			return "", err
		}
		return newTask.Id, nil
	}

	tasks, err = readFile()
	if err != nil {
		return "", err
	}

	tasks = append(tasks, newTask)
	err = createFile(tasks)
	if err != nil {
		return "", err
	}

	return newTask.Id, err
}

func (t *Task) UpdateDescription(newDescription string) {
	t.Description = newDescription
	t.UpdatedAt = time.Now()
}

func UpdateDescription(taskID, taskTitle string) error {
	tasks, err := ReadAllTasks()
	if err != nil {
		return err
	}

	taskFound := false
	for i := range tasks {
		if tasks[i].Id == taskID {
			tasks[i].UpdateDescription(taskTitle) // Use method to update task description
			taskFound = true
			break
		}
	}
	if !taskFound {
		return fmt.Errorf("task with id %s not found", taskID)
	}

	err = createFile(tasks)
	if err != nil {
		return err
	}
	return nil
}

func (t *Task) UpdateStatus(newStatus string) {
	t.Status = newStatus
	t.UpdatedAt = time.Now()
}

// updateTaskStatus updates a task's status by its ID.
func updateTaskStatus(taskID, status string) error {
	tasks, err := ReadAllTasks()
	if err != nil {
		return err
	}

	taskFound := false
	for i := range tasks {
		if tasks[i].Id == taskID {
			tasks[i].UpdateStatus(status)
			taskFound = true
			break
		}
	}

	if !taskFound {
		return fmt.Errorf("task with id %s not found", taskID)
	}

	return createFile(tasks)
}

// MarkDone marks a task as done.
func MarkDone(taskID string) error {
	return updateTaskStatus(taskID, "done")
}

// MarkInProgress marks a task as in-progress.
func MarkInProgress(taskID string) error {
	return updateTaskStatus(taskID, "in-progress")
}

// MarkToDo marks a task as to-do.
func MarkToDo(taskID string) error {
	return updateTaskStatus(taskID, "todo")
}

func ReadTasksStatus(status string) ([]Task, error) {
	tasks, err := ReadAllTasks()
	if err != nil {
		return nil, err
	}
	var tasksDone []Task
	var tasksInProgress []Task
	var tasksTodo []Task

	switch status {
	case "done":
		for _, task := range tasks {
			if task.Status == status {
				tasksDone = append(tasksDone, task)
			}
		}
		return tasksDone, nil
	case "in-progress":
		for _, task := range tasks {
			if task.Status == status {
				tasksInProgress = append(tasksInProgress, task)
			}
		}
		return tasksInProgress, nil
	case "todo":
		for _, task := range tasks {
			if task.Status == status {
				tasksTodo = append(tasksTodo, task)
			}
		}
		return tasksTodo, nil
	}
	return nil, fmt.Errorf("unknown status: %s\n", status)
}

func DeleteTask(taskID string) error {
	tasks, err := ReadAllTasks()
	if err != nil {
		return err
	}

	taskFound := false
	for i := range tasks {
		if tasks[i].Id == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			taskFound = true
			break
		}
	}
	if !taskFound {
		return fmt.Errorf("task with id %s not found", taskID)
	}

	err = createFile(tasks)
	if err != nil {
		return err
	}
	return nil
}
