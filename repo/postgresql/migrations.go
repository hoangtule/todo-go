package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Migrate(db *sql.DB) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			title TEXT,
			description TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP,
			completed_at TIMESTAMP
		);
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'todos' created or already exists")

	return nil
}
