package repository_clients

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"lab4/internal/models"
	"lab4/internal/repository/postgres"
)

type ClientsPostgres struct {
	db *sqlx.DB
}

func NewClientsPostgres(db *sqlx.DB) *ClientsPostgres {
	return &ClientsPostgres{db: db}
}

func (l *ClientsPostgres) CreateClient(client *models.Client) error {
	query := fmt.Sprintf("INSERT INTO %s (full_name,age,sex) values ($1,$2,$3) RETURNING id", postgres.ClientsTable)
	err := l.db.QueryRow(query, client.FullName, client.Age, client.Sex).Scan(&client.Id)
	return err
}

func (l *ClientsPostgres) GetAllClients() ([]models.Client, error) {
	data := make([]models.Client, 0)
	query := fmt.Sprintf("SELECT * FROM %s", postgres.ClientsTable)
	err := l.db.Select(&data, query)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (l *ClientsPostgres) UpdateClient(client *models.Client) error {
	query := fmt.Sprintf("UPDATE %s SET full_name=$2, age=$3, sex=$4 WHERE id=$1", postgres.ClientsTable)
	_, err := l.db.Exec(query, client.Id, client.FullName, client.Age, client.Sex)
	return err
}

func (l *ClientsPostgres) GetClientById(client *models.Client) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", postgres.ClientsTable)
	err := l.db.Get(client, query, client.Id)
	return err
}

func (l *ClientsPostgres) DeleteClient(client *models.Client) error {
	err := l.GetClientById(client)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postgres.ClientsTable)
	_, err = l.db.Exec(query, client.Id)
	return err
}
