package task

import (
	"context"
	"log/slog"
	"time"

	"task-planer-back/internal/models"

	"github.com/google/uuid"
)

type Service struct {
	Repo Storage
}

func NewServices(repo Storage) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) TaskServices(ctx context.Context, taskDTO *CreateTaskDTO) {

	t := &Task{
		ID:          uuid.New(),
		Name:        taskDTO.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: taskDTO.Description,
		Priority:    models.LowPriority,
		UserID:      taskDTO.UserID,
		IsCompleted: taskDTO.IsCompleted,
	}
	s.Repo.CreateTask(ctx, t)
	return
}

func (s *Service) CreateTask(ctx context.Context, dto *CreateTaskDTO) (*Task, error) {
	t := &Task{
		Name:        dto.Name,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Description: dto.Description,
		Priority:    models.LowPriority,
		UserID:      dto.UserID,
		IsCompleted: dto.IsCompleted,
	}
	task, err := s.Repo.CreateTask(ctx, t)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *Service) DeleteTask(ctx context.Context, id string) error {
	err := s.Repo.DeleteTask(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ChangeTask(ctx context.Context, dto ChangeTaskDTO) (*Task, error) {
	slog.Info("task ID", "ID", dto.ID)
	task, err := s.Repo.GetTask(ctx, dto.ID)
	if err != nil {
		slog.Error("get task error", "err", err)
		return nil, err
	}

	if dto.Name != nil {
		task.Name = *dto.Name
	}

	if dto.Description != nil {
		task.Description = *dto.Description
	}

	if dto.IsCompleted != nil {
		task.IsCompleted = *dto.IsCompleted
	}

	if dto.Priority != nil {
		task.Priority = *dto.Priority
	}

	updateTask, err := s.Repo.UpdateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	return updateTask, nil

}
