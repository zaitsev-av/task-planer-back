package task

import (
	"task-planer-back/pkg/client/postgresql"
)

type TaskRepository struct {
	DB *postgresql.Client
}

func (r *TaskRepository) Create(task *Task) error {
	return nil
}
