package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/anderson2093/ejercicios-go/4.-task-manager/config"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
}

func GetTasks() ([]Task, error) {
	rows, err := config.DB.Query("SELECT id, title, description, created_at, status FROM tasks")
	if err != nil {
		log.Println("❌ Error en la consulta de tareas:", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.Status); err != nil {
			log.Println("❌ Error al escanear fila:", err)
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		log.Println("❌ Error al recorrer las filas:", err)
		return nil, err
	}

	return tasks, nil
}

func CreateTask(title string, description string) error {
	if title == "" || description == "" {
		return fmt.Errorf("title y description son obligatorios")
	}
	_, err := config.DB.Exec("INSERT INTO tasks (title,description,status, created_at) VALUES ($1, $2, $3,$4)", title, description, "pending", time.Now())
	return err
}

func DeleteTask(id int) error {
	// Reusar la función para verificar si la tarea existe
	_, err := GetTaskByID(id)
	if err != nil {
		return err // Si la tarea no existe, se retorna el error directamente
	}

	// Si existe, proceder con la eliminación
	_, err = config.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error al eliminar la tarea: %v", err)
	}

	return nil
}

func GetTaskByID(id int) (*Task, error) {
	var task Task
	err := config.DB.QueryRow("SELECT id, title, description, created_at, status FROM tasks WHERE id = $1", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.CreatedAt, &task.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("la tarea con ID %d no existe", id)
		}
		return nil, fmt.Errorf("error al obtener la tarea: %v", err)
	}

	return &task, nil
}
