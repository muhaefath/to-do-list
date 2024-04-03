package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/muhaefath/to-do-list/internals/handlers"
	"github.com/muhaefath/to-do-list/internals/repository"
	"github.com/muhaefath/to-do-list/internals/repository/product"
)

func SetupRoutes(app *fiber.App, todoRepo repository.TodoListRepository,
	orderRepo repository.OrderRepository,
	productRepo product.ProductRepository) {

	todoHandler := handlers.NewTodoHandler(todoRepo)
	marketplaceHandler := handlers.NewMarketplaceHandler(productRepo, orderRepo)

	// Use CORS middleware
	app.Use(cors.New())

	// Define routes
	v1Group := app.Group("/v1")
	v1Group.Get("/todos", handlers.GetTodos)
	v1Group.Post("/todos", handlers.CreateTodo)
	v1Group.Get("/todos/:id", handlers.GetTodo)
	v1Group.Put("/todos/:id", handlers.UpdateTodo)
	v1Group.Delete("/todos/:id", handlers.DeleteTodo)

	v2Group := app.Group("/v2")
	v2Group.Get("/todos", todoHandler.GetAll)
	v2Group.Post("/todos", todoHandler.CreateTodo)
	v2Group.Get("/todos/:id", todoHandler.GetByID)
	v2Group.Put("/todos/:id", todoHandler.UpdateTodoByID)
	v2Group.Delete("/todos/:id", todoHandler.DeleteTodoByID)

	v2Group.Get("/products", marketplaceHandler.GetAllProduct)
	v2Group.Get("/orders", marketplaceHandler.GetAllOrder)
	v2Group.Post("/orders", marketplaceHandler.CreateOrder)
	v2Group.Put("/product/:id", marketplaceHandler.UpdateProduct)
	v2Group.Delete("/orders/:id", marketplaceHandler.CancelOrder)
}
