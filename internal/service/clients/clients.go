package servise_clients

import (
	"lab4/internal/models"
	"lab4/internal/repository"
)

type ClientService struct {
	repo repository.TodoClients
}

func NewClientService(repo repository.TodoClients) *ClientService {
	return &ClientService{repo: repo}
}

func (c *ClientService) CreateClient(client *models.Client) error {
	err := client.Validate()
	if err != nil {
		return err
	}
	err = c.repo.CreateClient(client)
	return err
}

func (c *ClientService) GetClientById(client *models.Client) error {
	return c.repo.GetClientById(client)
}
func (c *ClientService) GetAllClients() ([]models.Client, error) {
	return c.repo.GetAllClients()
}

func (c *ClientService) UpdateClient(client *models.Client) error {
	err := client.Validate()
	if err != nil {
		return err
	}
	return c.repo.UpdateClient(client)
}

func (c *ClientService) DeleteClient(client *models.Client) error {
	return c.repo.DeleteClient(client)
}
