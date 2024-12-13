package task

import "github.com/google/uuid"

type DTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsCompleted bool      `json:"is_completed"`
	UserId      uuid.UUID `json:"user_id"`
}
