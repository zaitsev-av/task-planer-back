package services

import (
	"context"
	"task-planer-back/internal/models"
	"task-planer-back/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func (s *TaskService) TaskServices(ctx context.Context, task *models.Task) {

}
