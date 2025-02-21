package models

import (
	"time"

	"github.com/anderson2093/ejercicios-go/4.-task-manager/config"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Completed   bool      `json:"completed"`
}

func GetTasks() ([]Task, error) {
	rows, err := config.DB.Query("SELECT id, title,description ,completed, created_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func CreateTask(title string, description string) error {
	_, err := config.DB.Exec("INSERT INTO tasks (title,description,completed, created_at) VALUES ($1, $2, $3,$4)", title, description, false, time.Now())
	return err
}
