package task

import "context"

type Storage interface {
	CreateTask(ctx context.Context, task *Task) (*Task, error)
	GetTask(ctx context.Context, id string) (Task, error)
	DeleteTask(ctx context.Context, id string) error
	RenameTask(ctx context.Context, id string, name string) error
	ChangeDescriptionTask(ctx context.Context, id string, description string) error
}
