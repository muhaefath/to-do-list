package migrations

import (
	"database/sql"
	"fmt"
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

	for idx, file := range files {
		if idx == len(files)-1 {
			continue
		}

		fmt.Println("file: ", file)

		if file.IsDir() {
			continue
		}

		migrationFilePath := filepath.Join(migrationsDir, file.Name())
		migrationFile, err := ioutil.ReadFile(migrationFilePath)
		if err != nil {
			fmt.Println("err2: ", err)

			return err
		}

		_, err = db.Exec(string(migrationFile))
		if err != nil {
			fmt.Println("err3: ", err)

			return err
		}

		log.Printf("Applied migration: %s\n", file.Name())
	}

	return nil
}
