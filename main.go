package main

import (
	"fmt"
	"log"

	// "github.com/your-username/create"
	// "github.com/your-username/read"
	// "github.com/your-username/update"
	// "github.com/your-username/delete"
)

func main() {
	taskListID := "your-task-list-id"

	// Create a new task
	err := create.CreateTask(taskListID, "Task Title")
	if err != nil {
		log.Fatalf("Error creating task: %v", err)
	}

	// Read tasks from the task list
	err = read.ReadTasks(taskListID)
	if err != nil {
		log.Fatalf("Error reading tasks: %v", err)
	}

	// Update a task title
	err = update.UpdateTask(taskListID, "task-id", "New Task Title")
	if err != nil {
		log.Fatalf("Error updating task: %v", err)
	}

	// Delete a task
	err = delete.DeleteTask(taskListID, "task-id")
	if err != nil {
		log.Fatalf("Error deleting task: %v", err)
	}

	fmt.Println("All operations completed successfully.")
}
