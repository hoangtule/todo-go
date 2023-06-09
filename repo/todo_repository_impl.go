package repo

import (
	"errors"
	"strconv"
	"time"

	"new-example/entity"
)

type TodoRepositoryImpl struct {
	todos []*entity.Todo
}

func (repo *TodoRepositoryImpl) InitDB() error {
	return nil
}

func (repo *TodoRepositoryImpl) GetByID(id string) (*entity.Todo, error) {
	for _, todo := range repo.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, errors.New("Todo not found")
}

func (repo *TodoRepositoryImpl) GetAll() ([]*entity.Todo, error) {
	return repo.todos, nil
}

func (repo *TodoRepositoryImpl) Create(todo *entity.Todo) (*entity.Todo, error) {
	todo.ID = strconv.Itoa(len(repo.todos) + 1)
	todo.CreatedAt = time.Now()
	repo.todos = append(repo.todos, todo)
	return todo, nil
}

func (repo *TodoRepositoryImpl) Update(id string, todo *entity.Todo) (*entity.Todo, error) {
	for i, t := range repo.todos {
		if t.ID == todo.ID {
			todo.UpdatedAt = time.Now()
			repo.todos[i] = todo
			return todo, nil
		}
	}
	return nil, errors.New("Todo not found")
}

func (repo *TodoRepositoryImpl) Delete(id string) error {
	for i, todo := range repo.todos {
		if todo.ID == id {
			repo.todos = append(repo.todos[:i], repo.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("Todo not found")
}

// func NewTodoRepository() TodoRepository {
// 	return &TodoRepositoryImpl{
// 		todos: make([]*entity.Todo, 0),
// 	}
// }
