package sqlite

import (
	"database/sql"
	"fmt"
)

func Init() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return nil, fmt.Errorf("could not open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not establish a connection with the database: %v", err)
	}

	return db, nil
}
