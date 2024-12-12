package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id          uuid.UUID     `json:"id"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Name        string        `json:"name"`
	Priority    PriorityModel `json:"priority"`
	IsCompleted bool          `json:"is_completed"`
	Description string        `json:"description"`
	UserId      uuid.UUID     `json:"user_id"`
}

type TaskDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	UserId      uuid.UUID `json:"user_id"`
}
