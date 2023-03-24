package sqlite

import (
	"database/sql"
	"errors"
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
	todo := &entity.Todo{}

	row := repo.db.QueryRow(`SELECT id, title, description, created_at, updated_at, completed_at FROM todos WHERE id = ?`, id)
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.CompletedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("todo not found: %v", err)
		}
		return nil, fmt.Errorf("could not get todo: %v", err)
	}

	return todo, nil
}

func (repo *TodoRepositoryImpl) GetAll() ([]*entity.Todo, error) {
	var todos []*entity.Todo

	rows, err := repo.db.Query(`SELECT id, title, description, created_at, updated_at, completed_at FROM todos`)
	if err != nil {
		return nil, fmt.Errorf("could not get todos: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		todo := &entity.Todo{}

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.CompletedAt)
		if err != nil {
			return nil, fmt.Errorf("could not get todo: %v", err)
		}

		todos = append(todos, todo)
	}

	return todos, nil
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
