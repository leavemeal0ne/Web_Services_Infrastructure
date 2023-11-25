package repository

import (
	"github.com/jmoiron/sqlx"
	"lab4/internal/models"
	repositoryclients "lab4/internal/repository/postgres/clients"
	repositorypositions "lab4/internal/repository/postgres/positions"
	repositoryworkers "lab4/internal/repository/postgres/workers"
)

type TodoClients interface {
	CreateClient(client *models.Client) error
	GetAllClients() ([]models.Client, error)
	UpdateClient(client *models.Client) error
	GetClientById(client *models.Client) error
	DeleteClient(client *models.Client) error
}

type TodoWorkers interface {
	CreateWorker(worker *models.Worker) error
	GetAllWorkers() ([]models.Worker, error)
	UpdateWorker(worker *models.Worker) error
	GetWorkerById(worker *models.Worker) error
	DeleteWorker(worker *models.Worker) error
}

type TodoPositions interface {
	CreatePosition(position *models.Position) error
	GetAllPositions() ([]models.Position, error)
	UpdatePosition(position *models.Position) error
	GetPositionById(position *models.Position) error
	DeletePosition(position *models.Position) error
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
	return &Repository{
		TodoClients:   repositoryclients.NewClientsPostgres(db),
		TodoWorkers:   repositoryworkers.NewWorkersPostgres(db),
		TodoPositions: repositorypositions.NewPositionsPostgres(db),
	}
}
