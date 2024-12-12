package task

import (
	"context"
	"github.com/google/uuid"
	"task-planer-back/internal/models"
	"time"
)

type TaskService struct {
	Repo *TaskRepository
}

func (s *TaskService) TaskServices(ctx context.Context, taskDTO *DTO) {

	task := &Task{
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
