package routes

import (
	"net/http"
	"strings"

	"github.com/anderson2093/ejercicios-go/4.-task-manager/handlers"
)

func RegisterRoutes() {
	// Manejar rutas generales (/tasks)
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTasksHandler(w, r) // Maneja GET /tasks
		case http.MethodPost:
			handlers.CreateTaskHandler(w, r) // Maneja POST /tasks
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})

	// Manejar tareas por ID (/tasks/{id})
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(r.URL.Path, "/")

		// Validar que la ruta tenga un ID (Ej: /tasks/1)
		if len(pathParts) < 3 || pathParts[2] == "" {
			http.Error(w, "ID es requerido", http.StatusBadRequest)
			return
		}

		// Manejar DELETE /tasks/{id}
		if r.Method == http.MethodDelete {
			handlers.DeleteTaskHandler(w, r)
			return
		}

		// Manejar GET /tasks/{id}
		if r.Method == http.MethodGet {
			handlers.GetTaskByIDHandler(w, r)
			return
		}

		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	})
}
