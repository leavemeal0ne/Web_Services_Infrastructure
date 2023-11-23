package service

import "lab4/internal/repository"

type TodoClients interface {
}

type TodoWorkers interface {
}

type TodoPositions interface {
}

type Service struct {
    TodoClients
    TodoWorkers
    TodoPositions
}

func NewService(r *repository.Repository) *Service {
    return &Service{}
}
