// main.go

package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/muhaefath/to-do-list/config"
	"github.com/muhaefath/to-do-list/database"
	"github.com/muhaefath/to-do-list/internals/repository"
	"github.com/muhaefath/to-do-list/router"
)

func main() {
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Access configuration values
	port := config.Port
	logLevel := config.LogLevel
	dbHost := config.Database.Host
	// Access other configuration values similarly...

	// Now you can use these values in your application
	log.Printf("Server is running on port %d\n", port)
	log.Printf("Log level: %s\n", logLevel)
	log.Printf("Database host: %s\n", dbHost)

	// Establish database connection
	db, err := database.ConnectDB(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Create a new Fiber instance
	app := fiber.New()

	// Define your routes here
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	todoRepo := repository.NewTodoListRepository(db)
	router.SetupRoutes(app, *todoRepo)

	// Start the server
	app.Listen(fmt.Sprintf(":%d", port))
}
