package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	"new-example/entity"
)

func CreateTodo(db *sql.DB, todo *entity.Todo) error {
	query := "INSERT INTO todos (id, title, description, created_at, updated_at, completed_at) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("could not prepare statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.ID, todo.Title, todo.Description, todo.CreatedAt, todo.UpdatedAt, todo.CompletedAt)
	if err != nil {
		return fmt.Errorf("could not create todo: %v", err)
	}

	return nil
}

func GetTodoByID(db *sql.DB, id string) (*entity.Todo, error) {
	todo := &entity.Todo{}
	row := db.QueryRow(`
		SELECT id, title, description, created_at, updated_at, completed_at
		FROM todos WHERE id = ?
	`, id)

	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt, &todo.CompletedAt)
	if err != nil {
		return nil, fmt.Errorf("could not get todo: %v", err)
	}

	return todo, nil
}

func UpdateTodoByID(db *sql.DB, id string, updates map[string]interface{}) error {
	var placeholders []string
	var values []interface{}
	for k, v := range updates {
		placeholders = append(placeholders, fmt.Sprintf("%s = ?", k))
		values = append(values, v)
	}

	query := fmt.Sprintf(`
		UPDATE todos SET %s WHERE id = ?
	`, strings.Join(placeholders, ", "))

	_, err := db.Exec(query, append(values, id)...)
	if err != nil {
		return fmt.Errorf("could not update todo: %v", err)
	}

	return nil
}

func DeleteTodoByID(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("could not delete todo: %v", err)
	}
	return nil
}

func GetTodos(db *sql.DB) ([]*entity.Todo, error) {
	var todos []*entity.Todo

	rows, err := db.Query(`
		SELECT id, title, description, created_at, updated_at, completed_at
		FROM todos
	`)
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
