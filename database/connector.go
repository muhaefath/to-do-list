package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/muhaefath/to-do-list/config"
)

func ConnectDB(config config.Config) (*sql.DB, error) {
	// Construct connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.Username,
		config.Database.Password, config.Database.Name)

	// Open connection to PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		return nil, err
	}

	// Check if connection is successful
	err = db.Ping()
	if err != nil {
		log.Printf("Error pinging database: %v\n", err)
		return nil, err
	}

	log.Println("Connected to database successfully")
	return db, nil
}
