package routes

import (
	"github.com/anderson2093/ejercicios-go/6.CleanArchitecture/internal/handler"
	"github.com/anderson2093/ejercicios-go/6.CleanArchitecture/internal/repository"
	"github.com/anderson2093/ejercicios-go/6.CleanArchitecture/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	taskRepo := repository.NewTaskRepository()
	taskUseCase := usecase.NewTaskUseCase(taskRepo)
	taskHandler := handler.NewTaskHandler(taskUseCase)

	r.Post("/tasks", taskHandler.CreateTask)
	r.Get("/tasks", taskHandler.GetAllTasks)
}
