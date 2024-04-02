package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhaefath/to-do-list/internals/models"
	"github.com/muhaefath/to-do-list/internals/repository"
)

type TodoHandler struct {
	todoRepo repository.TodoListRepository
}

func NewTodoHandler(todoRepo repository.TodoListRepository) *TodoHandler {
	return &TodoHandler{todoRepo}
}

func (h *TodoHandler) GetAll(c *fiber.Ctx) error {
	todos, err := h.todoRepo.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve todos: " + err.Error())
	}

	return c.JSON(todos)
}

func (h *TodoHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	todo, err := h.todoRepo.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to retrieve todos: " + err.Error())
	}

	return c.JSON(todo)
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	if err := h.todoRepo.CreateTodo(&todo); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(todo)
}

func (h *TodoHandler) UpdateTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedTodo models.Todo
	if err := c.BodyParser(&updatedTodo); err != nil {
		return err
	}

	if err := h.todoRepo.UpdateTodoByID(id, &updatedTodo); err != nil {
		return err
	}

	return c.JSON(updatedTodo)
}

func (h *TodoHandler) DeleteTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.todoRepo.DeleteTodoByID(id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
