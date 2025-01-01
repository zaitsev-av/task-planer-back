package task

import (
	"context"
	"errors"
	"fmt"

	"task-planer-back/internal/task"
	"task-planer-back/pkg/client/postgresql"

	"github.com/jackc/pgconn"
)

type Repository struct {
	db postgresql.Client
}

func (r *Repository) CreateTask(ctx context.Context, task *task.Task) (*task.Task, error) {
	q := `
		INSERT INTO tasks 
		(created_at, updated_at, name, priority, is_completed, description, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at, name, priority, is_completed, description, user_id
		`
	err := r.db.QueryRow(ctx, q,
		task.CreatedAt,
		task.UpdatedAt,
		task.Name,
		task.Priority,
		task.IsCompleted,
		task.Description,
		task.UserID,
	).Scan(
		&task.ID,
		&task.CreatedAt,
		&task.UpdatedAt,
		&task.Name,
		&task.Priority,
		&task.IsCompleted,
		&task.Description,
		&task.UserID,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr)
			return nil, pgErr
		}
		return nil, err
	}
	return task, nil
}
func (r *Repository) GetTask(ctx context.Context, id string) (task.Task, error) {
	panic("")
}

func (r *Repository) DeleteTask(ctx context.Context, id string) error {
	q := `
	DELETE FROM public.tasks WHERE id=$1;
	`
	_, err := r.db.Exec(ctx, q, id)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr)
			return err
		}
		return err
	}
	return nil
}

func (r *Repository) RenameTask(ctx context.Context, id string, name string) (*task.ChangeNameDTO, error) {
	var updTask task.ChangeNameDTO

	q := `
	UPDATE public.tasks 
	SET name = $2, 
	    updated_at = NOW()
	WHERE id = $1
	RETURNING id, name;
`

	err := r.db.QueryRow(ctx, q, id, name).Scan(&updTask.ID, &updTask.Name)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr)
			return nil, pgErr
		}
		return nil, err
	}
	return &updTask, nil
}

func (r *Repository) ChangeDescriptionTask(ctx context.Context, id string, description string) error {
	panic("")
}

func NewRepository(client postgresql.Client) task.Storage {
	return &Repository{
		db: client,
	}
}
