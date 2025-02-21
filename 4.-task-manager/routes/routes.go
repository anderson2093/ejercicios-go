package routes

import (
	"net/http"

	"github.com/anderson2093/ejercicios-go/4.-task-manager/handlers"
)

func RegisterRoutes() {
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

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) { // 🔥 Para manejar /tasks/{id}
		if r.Method == "DELETE" {
			handlers.DeleteTaskHandler(w, r)
		} else {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	})
}
