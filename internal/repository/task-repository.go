package repository

import (
	"database/sql"
	"task-planer-back/internal/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func (r *TaskRepository) Create(task *models.Task) error {

	return nil
}
