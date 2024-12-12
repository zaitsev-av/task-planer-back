package services

import (
	"context"
	"github.com/google/uuid"
	"task-planer-back/internal/models"
	"task-planer-back/internal/repository"
	"time"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func (s *TaskService) TaskServices(ctx context.Context, taskDTO *models.TaskDTO) {

	task := &models.Task{
		Id:          uuid.New(),
		Name:        taskDTO.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: taskDTO.Description,
		Priority:    models.LowPriority,
		UserId:      taskDTO.UserId,
		IsCompleted: taskDTO.IsCompleted,
	}

	s.Repo.Create(task)
	return
}
