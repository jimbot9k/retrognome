package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
