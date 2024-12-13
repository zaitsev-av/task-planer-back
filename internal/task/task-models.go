package task

import (
	"github.com/google/uuid"
	"task-planer-back/internal/models"
	"time"
)

type Task struct {
	Id          uuid.UUID            `json:"id"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Name        string               `json:"name"`
	Priority    models.PriorityModel `json:"priority"`
	IsCompleted bool                 `json:"is_completed"`
	Description string               `json:"description"`
	UserId      uuid.UUID            `json:"user_id"`
}
