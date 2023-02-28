package service

import (
	"GOLANG-TODOS/errs"
	"GOLANG-TODOS/logs"
	"GOLANG-TODOS/repository"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type todoService struct {
	todoRepo repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{todoRepo: todoRepo}
}

func (s *todoService) NewTodo(request TodoRequest) (*TodoResponse, error) {
	todo := repository.Todo{
		Title: request.Title,
	}
	newTodo, err := s.todoRepo.Create(todo)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	response := TodoResponse{
		ID:        newTodo.ID,
		Title:     newTodo.Title,
		Status:    newTodo.Status,
		UpdatedAt: newTodo.UpdatedAt,
	}

	return &response, nil
}

func (s *todoService) UpdateTodo(request TodoUpdateRequest) (*TodoResponse, error) {

	todo, err := s.todoRepo.GetByID(int(request.ID))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		logs.Error(err)
		return nil, errs.NewNotFoundError(fmt.Sprintf("Todo with id %d not found", request.ID))
	}

	todo.Title = request.Title
	todo.Status = request.Status
	todoSave, err := s.todoRepo.Update(*todo)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := TodoResponse{
		ID:        todoSave.ID,
		Title:     todoSave.Title,
		Status:    todoSave.Status,
		UpdatedAt: todoSave.UpdatedAt,
	}
	return &response, nil
}

func (s *todoService) GetTodos() ([]TodoResponse, error) {
	todos, err := s.todoRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	var todoResponses []TodoResponse
	for _, todo := range todos {
		todoResponses = append(todoResponses, TodoResponse{
			ID:        todo.ID,
			Title:     todo.Title,
			Status:    todo.Status,
			UpdatedAt: todo.UpdatedAt,
		})
	}
	return todoResponses, nil
}

func (s *todoService) GetTodo(id int) (*TodoResponse, error) {
	todo, err := s.todoRepo.GetByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		logs.Error(err)
		return nil, errs.NewNotFoundError(fmt.Sprintf("Todo with id %d not found", id))
	}

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	todoResponse := &TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		Status:    todo.Status,
		UpdatedAt: todo.UpdatedAt,
	}
	return todoResponse, nil
}

func (s *todoService) DeleteTodo(id uint) error {
	err := s.todoRepo.Delete(id)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}
