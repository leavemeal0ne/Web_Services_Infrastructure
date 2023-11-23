package repository

import "github.com/jmoiron/sqlx"

type TodoClients interface {
}

type TodoWorkers interface {
}

type TodoPositions interface {
}

type Repository struct {
    TodoClients
    TodoWorkers
    TodoPositions
}

func NewRepository(db *sqlx.DB) *Repository {
    return &Repository{}
}

func NewPostgresRepository(db *sqlx.DB) *Repository {
    return &Repository{}
}
