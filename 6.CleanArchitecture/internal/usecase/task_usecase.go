package usecase

import (
	"github.com/anderson2093/ejercicios-go/6.CleanArchitecture/internal/domain"
	"github.com/anderson2093/ejercicios-go/6.CleanArchitecture/internal/repository"
)

type TaskUseCase struct {
	repo *repository.TaskRepository
}

func NewTaskUseCase(repo *repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo}
}

func (uc *TaskUseCase) CreateTask(task domain.Task) error {
	return uc.repo.CreateTask(task)
}

func (uc *TaskUseCase) GetAllTasks() ([]domain.Task, error) {
	return uc.repo.GetAllTasks()
}
