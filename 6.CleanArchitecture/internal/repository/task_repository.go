package repository

import (
	"context"
	"log"

	"github.com/anderson2093/ejercicios-go/6.CleanArchitecture/config"
	"github.com/anderson2093/ejercicios-go/6.CleanArchitecture/internal/domain"
)

type TaskRepository struct{}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

func (r *TaskRepository) CreateTask(task domain.Task) error {
	query := "INSERT INTO tasks (title, description) VALUES ($1, $2)"
	_, err := config.DB.Exec(context.Background(), query, task.Title, task.Description)
	if err != nil {
		log.Println("❌ Error creando tarea:", err)
		return err
	}
	return nil
}

func (r *TaskRepository) GetAllTasks() ([]domain.Task, error) {
	query := "SELECT id, title, description FROM tasks"
	rows, err := config.DB.Query(context.Background(), query)
	if err != nil {
		log.Println("❌ Error obteniendo tareas:", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
