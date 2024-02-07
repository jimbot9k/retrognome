package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteDB() *sql.DB {
	// Implement logic to open a connection to a SQLite database
	// Use the database/sql package to open a connection
	db, err := sql.Open("sqlite3", "file:retrognome.sqlite?mode=rwc")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func UpdateSchema(db *sql.DB) error {
	files, err := os.ReadDir("internal/sql/")
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	for _, file := range fileNames {
		fileData, err := os.ReadFile(fmt.Sprintf("internal/sql/%s", file))
		if err != nil {
			log.Fatal(err)
		}

		/// Check if migrations table, if not create it
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS migrations (id INTEGER PRIMARY KEY, name TEXT, created_at INTEGER)")
		if err != nil {
			log.Fatal(err)
			return err
		}

		/// Check if migration has been applied
		if db.QueryRow("SELECT name FROM migrations WHERE name = ?", file).Scan(new(string)) == nil {
			log.Printf("Migration %s has already been applied", file)
			continue
		}

		_, err = db.Exec(string(fileData))
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Printf("Migration %s applied successfully", file)

		_, err = db.Exec("INSERT INTO migrations (name, created_at) VALUES (?, ?)", file, time.Now().Unix())
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	log.Printf("Schema updated successfully")
	return nil
}
