package repository_workers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"lab4/internal/models"
	"lab4/internal/repository/postgres"
)

type WorkersPostgres struct {
	db *sqlx.DB
}

func NewWorkersPostgres(db *sqlx.DB) *WorkersPostgres {
	return &WorkersPostgres{db: db}
}

func (l *WorkersPostgres) CreateWorker(worker *models.Worker) error {
	query := fmt.Sprintf("INSERT INTO %s (full_name,age,sex,position_id) values ($1,$2,$3,$4) RETURNING id", postgres.WorkersTable)
	err := l.db.QueryRow(query, worker.FullName, worker.Age, worker.Sex, worker.PositionId).Scan(&worker.Id)
	return err
}

func (l *WorkersPostgres) GetAllWorkers() ([]models.Worker, error) {
	data := make([]models.Worker, 0)
	query := fmt.Sprintf("SELECT * FROM %s", postgres.WorkersTable)
	err := l.db.Select(&data, query)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (l *WorkersPostgres) UpdateWorker(worker *models.Worker) error {
	query := fmt.Sprintf("UPDATE %s SET full_name=$2, age=$3, sex=$4, position_id=$5 WHERE id=$1", postgres.WorkersTable)
	_, err := l.db.Exec(query, worker.Id, worker.FullName, worker.Age, worker.Sex, worker.PositionId)
	return err
}

func (l *WorkersPostgres) GetWorkerById(worker *models.Worker) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", postgres.WorkersTable)
	err := l.db.Get(worker, query, worker.Id)
	return err
}

func (l *WorkersPostgres) DeleteWorker(worker *models.Worker) error {
	err := l.GetWorkerById(worker)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.WorkersTable)
	_, err = l.db.Exec(query, worker.Id)
	return err
}
