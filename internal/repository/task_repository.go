package repository

import "database/sql"

type TaskRepository struct {
	DB *sql.DB
}

func (r *TaskRepository) Create() {

}
