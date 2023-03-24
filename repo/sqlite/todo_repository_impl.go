package sqlite

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"new-example/entity"
)

type TodoRepositoryImpl struct {
	db *sql.DB
}

func (repo *TodoRepositoryImpl) InitDB() error {
	db, err := Init("./database.db")
	if err != nil {
		return err
	}

	err = Migrate(db)
	if err != nil {
		return err
	}

	repo.db = db
	return nil
}

func (repo *TodoRepositoryImpl) GetByID(id string) (*entity.Todo, error) {
	return GetTodoByID(repo.db, id)
}

func (repo *TodoRepositoryImpl) GetAll() ([]*entity.Todo, error) {
	return GetTodos(repo.db)
}

func GenerateRandomID() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	randNum := rand.Int63n(1000)
	id := timestamp*1000 + randNum
	return strconv.FormatInt(id, 10)
}

func (repo *TodoRepositoryImpl) Create(todo *entity.Todo) (*entity.Todo, error) {
	todo.ID = GenerateRandomID()
	todo.CreatedAt = time.Now()

	err := CreateTodo(repo.db, todo)
	if err != nil {
		return nil, fmt.Errorf("could not create todo: %v", err)
	}

	return todo, nil
}

func (repo *TodoRepositoryImpl) Update(id string, todo *entity.Todo) (*entity.Todo, error) {
	updates := make(map[string]interface{})
	if todo.Title != "" {
		updates["title"] = todo.Title
	}
	if todo.Description != "" {
		updates["description"] = todo.Description
	}
	if !todo.UpdatedAt.IsZero() {
		updates["updated_at"] = todo.UpdatedAt
	}
	if !todo.CompletedAt.IsZero() {
		updates["completed_at"] = todo.CompletedAt
	}

	err := UpdateTodoByID(repo.db, id, updates)
	if err != nil {
		return nil, fmt.Errorf("could not update todo: %v", err)
	}

	return todo, nil
}

func (repo *TodoRepositoryImpl) Delete(id string) error {
	err := DeleteTodoByID(repo.db, id)
	if err != nil {
		return err
	}
	return nil
}
