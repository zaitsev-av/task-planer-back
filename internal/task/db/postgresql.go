package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"task-planer-back/internal/task"
	"task-planer-back/pkg/client/postgresql"
)

type Repository struct {
	db postgresql.Client
}

func (r *Repository) CreateTask(ctx context.Context, task *task.Task) error {
	q := `
		INSERT INTO tasks 
		(created_at, updated_at, name, priority, is_completed, description, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING created_at, updated_at, name, priority, is_completed, description, user_id
		`
	err := r.db.QueryRow(ctx, q,
		task.CreatedAt,   // $1
		task.UpdatedAt,   // $2
		task.Name,        // $3
		task.Priority,    // $4
		task.IsCompleted, // $5
		task.Description, // $6
		task.UserId,      // $7
	).Scan(
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.Name,
		&task.Priority,
		&task.IsCompleted,
		&task.Description,
		&task.UserId,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr)
			return pgErr
		}
		return err
	}
	return nil
}
func (r *Repository) GetTask(ctx context.Context, id string) (task.Task, error) {
	panic("")
}

func (r *Repository) DeleteTask(ctx context.Context, is string) error {
	panic("")
}

func (r *Repository) RenameTask(ctx context.Context, id string, name string) error {
	panic("")
}

func (r *Repository) ChangeDescriptionTask(ctx context.Context, id string, description string) error {
	panic("")
}

func NewRepository(client postgresql.Client) task.Storage {
	return &Repository{
		db: client,
	}
}
