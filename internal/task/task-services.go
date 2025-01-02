package task

import (
	"context"
	"fmt"
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

func (s *Service) ChangeTask(ctx context.Context, id, name string) (*ChangeNameDTO, error) {

	res, err := s.Repo.RenameTask(ctx, id, name)
	if err != nil {
		fmt.Println(err, "change task error")
		return nil, err
	}

	return res, nil

}
