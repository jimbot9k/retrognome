package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func UpdateSchema(db *sql.DB) error {

	const sqlPath = "internal/database/sql/"

	files, err := os.ReadDir(sqlPath)
	if err != nil {
		log.Fatal(err)
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	for _, file := range fileNames {
		fileData, err := os.ReadFile(fmt.Sprintf("%s%s", sqlPath, file))
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
