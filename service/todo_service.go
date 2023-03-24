package service

import (
	"errors"

	"new-example/entity"
)

type TodoRepository interface {
	InitDB() error
	GetByID(id string) (*entity.Todo, error)
	GetAll() ([]*entity.Todo, error)
	Create(todo *entity.Todo) (*entity.Todo, error)
	Update(id string, todo *entity.Todo) (*entity.Todo, error)
	Delete(id string) error
}

type TodoService struct {
	TodoRepo TodoRepository
}

func (s *TodoService) InitDB() error {
	return nil
}

func (s *TodoService) GetAll() ([]*entity.Todo, error) {
	return s.TodoRepo.GetAll()
}

func (s *TodoService) GetByID(id string) (*entity.Todo, error) {
	//... biz
	//...
	return s.TodoRepo.GetByID(id)
}

func (s *TodoService) Create(todo *entity.Todo) (*entity.Todo, error) {
	if todo.Title == "" {
		return nil, errors.New("Title cannot be empty")
	}
	return s.TodoRepo.Create(todo)
}

func (s *TodoService) Update(id string, todo *entity.Todo) (*entity.Todo, error) {
	if todo.Title == "" {
		return nil, errors.New("Title cannot be empty")
	}
	return s.TodoRepo.Update(id, todo)
}

func (s *TodoService) Delete(id string) error {
	return s.TodoRepo.Delete(id)
}
