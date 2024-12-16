package task

import (
	"github.com/google/uuid"
	"task-planer-back/internal/models"
)

type CreateTaskDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	UserId      uuid.UUID `json:"user_id"`
}

type ChangeNameDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ChangeDescriptionDTO struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type ChangeIsCompletedDTO struct {
	Id          string `json:"id"`
	IsCompleted string `json:"is_completed"`
}

type ChangePriorityDTO struct {
	Id       string               `json:"id"`
	Priority models.PriorityModel `json:"priority"`
}
