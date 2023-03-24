package db

import (
	"database/sql"
	"fmt"
)

func Migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id TEXT PRIMARY KEY,
			title TEXT,
			description TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			completed_at TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("could not create todos table: %v", err)
	}

	return nil
}
