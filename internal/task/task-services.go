package task

import (
	"context"
	"github.com/google/uuid"
	"task-planer-back/internal/models"
	"time"
)

type Service struct {
	Repo Storage
}

func (s *Service) TaskServices(ctx context.Context, taskDTO *DTO) {

	t := &Task{
		Id:          uuid.New(),
		Name:        taskDTO.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: taskDTO.Description,
		Priority:    models.LowPriority,
		UserId:      taskDTO.UserId,
		IsCompleted: taskDTO.IsCompleted,
	}
	s.Repo.CreateTask(ctx, t)
	return
}

func (s *Service) CreateTask(ctx context.Context, dto DTO) error {
	t := &Task{
		Name:        dto.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: dto.Description,
		Priority:    models.LowPriority,
		UserId:      dto.UserId,
		IsCompleted: dto.IsCompleted,
	}
	err := s.Repo.CreateTask(ctx, t)
	if err != nil {
		return err
	}
	//s.Repo.Create(task)
	return nil
}

func NewServices(repo Storage) *Service {
	return &Service{
		Repo: repo,
	}
}
