package routes

import (
	"github.com/anderson2093/ejercicios-go/5.chi-task/handlers"
	"github.com/go-chi/chi/v5"
)

// Configurar rutas
func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/tasks", handlers.GetTasksHandler)
	r.Post("/tasks", handlers.CreateTaskHandler)
	return r
}
