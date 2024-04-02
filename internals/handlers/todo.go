package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhaefath/to-do-list/internals/models"
	"github.com/muhaefath/to-do-list/internals/repository"
)

func GetTodos(c *fiber.Ctx) error {
	todos, err := repository.GetAllTodos()
	if err != nil {
		return err
	}
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return err
	}
	err := repository.CreateTodo(&todo)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(todo)
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo, err := repository.GetTodoByID(id)
	if err != nil {
		return err
	}
	return c.JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return err
	}
	updatedTodo, err := repository.UpdateTodoByID(id, &todo)
	if err != nil {
		return err
	}
	return c.JSON(updatedTodo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	err := repository.DeleteTodoByID(id)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusNoContent)
}
