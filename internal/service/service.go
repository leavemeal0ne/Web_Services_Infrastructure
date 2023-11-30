package service

import (
    "lab4/internal/models"
    "lab4/internal/repository"
    servise_clients "lab4/internal/service/clients"
    servicelogs "lab4/internal/service/logs"
    servise_positions "lab4/internal/service/positions"
    servise_workers "lab4/internal/service/workers"
)

type TodoClients interface {
    CreateClient(client *models.Client) error
    GetClientById(client *models.Client) error
    GetAllClients() ([]models.Client, error)
    UpdateClient(client *models.Client) error
    DeleteClient(client *models.Client) error
}

type TodoWorkers interface {
    CreateWorker(worker *models.Worker) error
    GetWorkerById(worker *models.Worker) error
    GetAllWorkers() ([]models.Worker, error)
    UpdateWorker(worker *models.Worker) error
    DeleteWorker(worker *models.Worker) error
}

type TodoPositions interface {
    CreatePosition(position *models.Position) error
    GetPositionById(position *models.Position) error
    GetAllPositions() ([]models.Position, error)
    UpdatePosition(position *models.Position) error
    DeletePosition(position *models.Position) error
}

type MongoDBLog interface {
    InsertLog(data models.LogData) error
}

type Service struct {
    TodoClients
    TodoWorkers
    TodoPositions
    MongoDBLog
}

func NewService(r *repository.Repository) *Service {
    return &Service{
        TodoClients:   servise_clients.NewClientService(r),
        TodoWorkers:   servise_workers.NewWorkerService(r),
        TodoPositions: servise_positions.NewPositionService(r),
        MongoDBLog:    servicelogs.NewLogService(r),
    }
}
