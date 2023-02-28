package service

import (
	"time"
)

type TodoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Status    bool      `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoRequest struct {
	Title string `json:"title" validate:"required"`
}

type TodoUpdateRequest struct {
	ID     uint   `json:"id" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Status bool   `json:"status" validate:"required"`
}

type TodoService interface {
	GetTodos() ([]TodoResponse, error)
	GetTodo(int) (*TodoResponse, error)
	NewTodo(TodoRequest) (*TodoResponse, error)
	UpdateTodo(TodoUpdateRequest) (*TodoResponse, error)
	DeleteTodo(uint) error
}
