package repository

import (
	"errors"

	"gorm.io/gorm"
)

type todoRepositoryDB struct {
	db *gorm.DB
}

func NewTodoRepositoryDB(db *gorm.DB) TodoRepository {
	return &todoRepositoryDB{db: db}
}

func (r *todoRepositoryDB) GetAll() ([]Todo, error) {
	var todos []Todo
	err := r.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepositoryDB) GetByID(id int) (*Todo, error) {
	var todo Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepositoryDB) Create(todo Todo) (*Todo, error) {
	// todo = Todo{}
	err := r.db.Create(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepositoryDB) Update(todo Todo) (*Todo, error) {
	err := r.db.Save(&todo).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepositoryDB) Delete(id uint) error {
	tx := r.db.Unscoped().Delete(&Todo{}, id)
	if tx.Error != nil {
		return tx.Error
	} else if tx.RowsAffected < 1 {
		return errors.New("Todo not found")
	}
	return nil
}
