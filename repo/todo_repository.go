package repo

import (
	"new-example/entity"
)

type TodoRepository interface {
	GetByID(id string) (*entity.Todo, error)
	GetAll() ([]*entity.Todo, error)
	Create(todo *entity.Todo) (*entity.Todo, error)
	Update(id string, todo *entity.Todo) (*entity.Todo, error)
	Delete(id string) error
}
