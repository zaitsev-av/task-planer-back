package models

type Task struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Name        string `json:"name"`
	Priority    int    `json:"priority"`
	IsCompleted bool   `json:"is_completed"`
	User        string `json:"user"`
	UserId      string `json:"user_id"`
}
