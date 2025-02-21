package models

import (
	"context"

	"github.com/anderson2093/ejercicios-go/5.chi-task/config"
)

// Task representa una tarea en la BD
type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Obtener todas las tareas
func GetTasks() ([]Task, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, title, description FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

// Crear una nueva tarea
func CreateTask(title string, description string) error {
	_, err := config.DB.Exec(context.Background(), "INSERT INTO tasks (title, description) VALUES ($1, $2)", title, description)
	return err
}
