package task

import (
	"time"

	"task-planer-back/internal/models"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID            `json:"id"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Name        string               `json:"name"`
	Priority    models.PriorityModel `json:"priority"`
	IsCompleted bool                 `json:"is_completed"`
	Description string               `json:"description"`
	UserID      uuid.UUID            `json:"user_id"`
}
