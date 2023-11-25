package servise_workers

import (
	"lab4/internal/models"
	"lab4/internal/repository"
)

type WorkerService struct {
	repo repository.TodoWorkers
}

func NewWorkerService(repo repository.TodoWorkers) *WorkerService {
	return &WorkerService{repo: repo}
}

func (c *WorkerService) CreateWorker(worker *models.Worker) error {
	err := worker.Validate()
	if err != nil {
		return err
	}
	err = c.repo.CreateWorker(worker)
	return err
}

func (c *WorkerService) GetWorkerById(worker *models.Worker) error {
	return c.repo.GetWorkerById(worker)
}
func (c *WorkerService) GetAllWorkers() ([]models.Worker, error) {
	return c.repo.GetAllWorkers()
}

func (c *WorkerService) UpdateWorker(worker *models.Worker) error {
	err := worker.Validate()
	if err != nil {
		return err
	}
	return c.repo.UpdateWorker(worker)
}

func (c *WorkerService) DeleteWorker(worker *models.Worker) error {
	return c.repo.DeleteWorker(worker)
}
