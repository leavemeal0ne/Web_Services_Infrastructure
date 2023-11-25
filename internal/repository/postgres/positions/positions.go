package repository_positions

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"lab4/internal/models"
	"lab4/internal/repository/postgres"
)

type PositionsPostgres struct {
	db *sqlx.DB
}

func NewPositionsPostgres(db *sqlx.DB) *PositionsPostgres {
	return &PositionsPostgres{db: db}
}

func (l *PositionsPostgres) CreatePosition(position *models.Position) error {
	query := fmt.Sprintf("INSERT INTO %s (title,salary,description) values ($1,$2,$3) RETURNING id", postgres.PositionsTable)
	err := l.db.QueryRow(query, position.Title, position.Salary, position.Description).Scan(&position.Id)
	return err
}

func (l *PositionsPostgres) GetAllPositions() ([]models.Position, error) {
	data := make([]models.Position, 0)
	query := fmt.Sprintf("SELECT * FROM %s", postgres.PositionsTable)
	err := l.db.Select(&data, query)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (l *PositionsPostgres) UpdatePosition(position *models.Position) error {
	query := fmt.Sprintf("UPDATE %s SET title=$2, salary=$3, description=$4 WHERE id=$1", postgres.PositionsTable)
	_, err := l.db.Exec(query, position.Id, position.Title, position.Salary, position.Description)
	return err
}

func (l *PositionsPostgres) GetPositionById(position *models.Position) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", postgres.PositionsTable)
	err := l.db.Get(position, query, position.Id)
	return err
}

func (l *PositionsPostgres) DeletePosition(position *models.Position) error {
	err := l.GetPositionById(position)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.PositionsTable)
	_, err = l.db.Exec(query, position.Id)
	return err
}
