package servise_positions

import (
	"lab4/internal/models"
	"lab4/internal/repository"
)

type PositionService struct {
	repo repository.TodoPositions
}

func NewPositionService(repo repository.TodoPositions) *PositionService {
	return &PositionService{repo: repo}
}

func (c *PositionService) CreatePosition(position *models.Position) error {
	err := position.Validate()
	if err != nil {
		return err
	}
	err = c.repo.CreatePosition(position)
	return err
}

func (c *PositionService) GetPositionById(position *models.Position) error {
	return c.repo.GetPositionById(position)
}
func (c *PositionService) GetAllPositions() ([]models.Position, error) {
	return c.repo.GetAllPositions()
}

func (c *PositionService) UpdatePosition(position *models.Position) error {
	err := position.Validate()
	if err != nil {
		return err
	}
	return c.repo.UpdatePosition(position)
}

func (c *PositionService) DeletePosition(position *models.Position) error {
	return c.repo.DeletePosition(position)
}
