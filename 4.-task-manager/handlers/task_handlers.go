package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/anderson2093/ejercicios-go/4.-task-manager/models"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetTasks()
	if err != nil {
		http.Error(w, "Erro al obtener tareas", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Error al leer datos", http.StatusBadRequest)
		return
	}

	// Validar que title y description no sean vacíos
	if task.Title == "" || task.Description == "" {
		http.Error(w, "Title y Description son obligatorios", http.StatusBadRequest)
		return
	}

	err := models.CreateTask(task.Title, task.Description)
	if err != nil {
		http.Error(w, "Error al crear tarea", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer ID de la URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "ID es requerido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Intentar eliminar la tarea
	err = models.DeleteTask(id)
	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al eliminar la tarea", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

func GetTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer ID de la URL (Ejemplo: /tasks/1)
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.Error(w, "ID es requerido", http.StatusBadRequest)
		return
	}

	// Convertir el ID a entero
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Buscar la tarea por ID
	task, err := models.GetTaskByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "no existe") {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Error al obtener la tarea", http.StatusInternalServerError)
		}
		return
	}

	// Responder con la tarea en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
