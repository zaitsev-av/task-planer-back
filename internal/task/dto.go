package task

import (
	"task-planer-back/internal/models"

	"github.com/google/uuid"
)

type CreateTaskDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	UserID      uuid.UUID `json:"user_id"`
}

type ChangeTaskDTO struct {
	ID          string                `json:"id"`
	Name        *string               `json:"name"`
	Description *string               `json:"description"`
	IsCompleted *string               `json:"is_completed"`
	Priority    *models.PriorityModel `json:"priority"`
}

type ChangeNameDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ChangeDescriptionDTO struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type ChangeIsCompletedDTO struct {
	ID          string `json:"id"`
	IsCompleted string `json:"is_completed"`
}

type ChangePriorityDTO struct {
	ID       string               `json:"id"`
	Priority models.PriorityModel `json:"priority"`
}
