package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anderson2093/ejercicios-go/5.chi-task/models"
)

// Obtener todas las tareas
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetTasks()
	if err != nil {
		http.Error(w, "Error al obtener tareas", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

// Crear una nueva tarea
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Error al leer datos", http.StatusBadRequest)
		return
	}

	if err := models.CreateTask(task.Title, task.Description); err != nil {
		http.Error(w, "Error al crear tarea", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
