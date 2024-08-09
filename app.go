package main

import (
	"context"
	"fmt"
)

// Task represents a task in the application
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// CreateTask creates a new task
func (a *App) CreateTask(name string) error {
	// Implementation of task creation
	fmt.Printf("Task created: %s\n", name)
	return nil
}

// GetTasks retrieves the list of tasks
func (a *App) GetTasks() ([]Task, error) {
	// Implementation to retrieve tasks
	return []Task{
		{ID: 1, Name: "Sample Task"},
	}, nil
}

// UpdateTask updates a task by ID
func (a *App) UpdateTask(id int, name string) error {
	// Implementation to update task
	fmt.Printf("Task updated: %d, %s\n", id, name)
	return nil
}

// DeleteTask deletes a task by ID
func (a *App) DeleteTask(id int) error {
	// Implementation to delete task
	fmt.Printf("Task deleted: %d\n", id)
	return nil
}
