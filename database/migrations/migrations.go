package migrations

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Migrate runs SQL migration files located in a given directory
func Migrate(db *sql.DB, migrationsDir string) error {
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		migrationFilePath := filepath.Join(migrationsDir, file.Name())
		migrationFile, err := ioutil.ReadFile(migrationFilePath)
		if err != nil {
			return err
		}

		_, err = db.Exec(string(migrationFile))
		if err != nil {
			return err
		}

		log.Printf("Applied migration: %s\n", file.Name())
	}

	return nil
}
