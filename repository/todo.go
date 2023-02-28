package repository

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title  string
	Status bool
}

type TodoRepository interface {
	GetAll() ([]Todo, error)
	GetByID(int) (*Todo, error)
	Create(Todo) (*Todo, error)
	Update(Todo) (*Todo, error)
	Delete(uint) error
}
