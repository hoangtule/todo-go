package service

import (
	"errors"

	"new-example/entity"
	"new-example/repo"
)

type TodoService interface {
	GetAll() ([]*entity.Todo, error)
	GetByID(id string) (*entity.Todo, error)
	Create(todo *entity.Todo) (*entity.Todo, error)
	Update(id string, todo *entity.Todo) (*entity.Todo, error)
	Delete(id string) error
}

type todoService struct {
	todoRepo repo.TodoRepository
}

func NewTodoService(todoRepo repo.TodoRepository) TodoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}

func (s *todoService) GetAll() ([]*entity.Todo, error) {
	return s.todoRepo.GetAll()
}

func (s *todoService) GetByID(id string) (*entity.Todo, error) {
	return s.todoRepo.GetByID(id)
}

func (s *todoService) Create(todo *entity.Todo) (*entity.Todo, error) {
	if todo.Title == "" {
		return nil, errors.New("Title cannot be empty")
	}
	return s.todoRepo.Create(todo)
}

func (s *todoService) Update(id string, todo *entity.Todo) (*entity.Todo, error) {
	if todo.Title == "" {
		return nil, errors.New("Title cannot be empty")
	}
	return s.todoRepo.Update(id, todo)
}

func (s *todoService) Delete(id string) error {
	return s.todoRepo.Delete(id)
}
