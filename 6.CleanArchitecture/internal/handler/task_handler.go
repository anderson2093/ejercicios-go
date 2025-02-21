package handler

import (
	"encoding/json"
	"net/http"

	"github.com/anderson2093/ejercicios-go/6.CleanArchitecture/internal/domain"
	"github.com/anderson2093/ejercicios-go/6.CleanArchitecture/internal/usecase"
)

type TaskHandler struct {
	usecase *usecase.TaskUseCase
}

func NewTaskHandler(usecase *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{usecase}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "❌ Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	if err := h.usecase.CreateTask(task); err != nil {
		http.Error(w, "❌ No se pudo crear la tarea", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Tarea creada con éxito"})
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.usecase.GetAllTasks()
	if err != nil {
		http.Error(w, "❌ No se pudieron obtener las tareas", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}
