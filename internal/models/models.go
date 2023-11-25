package models

import "github.com/go-playground/validator/v10"

type Worker struct {
	Id         int    `json:"id" db:"id"`
	FullName   string `json:"full_name" db:"full_name" validate:"required,min=6,max=100"`
	Age        int    `json:"age" db:"age" validate:"required,gte=18,lte=120"`
	Sex        string `json:"sex" db:"sex" validate:"required,eq=female|eq=male"`
	PositionId int    `json:"position_id" db:"position_id" validate:"omitempty"`
}

func (w *Worker) Validate() error {
	return validator.New(validator.WithRequiredStructEnabled()).Struct(w)
}

type Client struct {
	Id       int    `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name" validate:"required,min=6,max=100"`
	Age      int    `json:"age" db:"age" validate:"required,gte=18,lte=120"`
	Sex      string `json:"sex" db:"sex" validate:"required,eq=female|eq=male"`
}

func (w *Client) Validate() error {
	return validator.New(validator.WithRequiredStructEnabled()).Struct(w)
}

type Position struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" validate:"required,min=6,max=100"`
	Salary      int    `json:"salary" db:"salary" validate:"required,gt=0"`
	Description string `json:"description" db:"description" validate:"omitempty,min=6,max=1000"`
}

func (w *Position) Validate() error {
	return validator.New(validator.WithRequiredStructEnabled()).Struct(w)
}
