package repository

import "github.com/muhaefath/to-do-list/internals/models"

// Example in-memory storage for simplicity
var todos []models.Todo

func GetAllTodos() ([]models.Todo, error) {
	return todos, nil
}

func CreateTodo(todo *models.Todo) error {
	todos = append(todos, *todo)
	return nil
}

func GetTodoByID(id string) (*models.Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			return &todo, nil
		}
	}
	return nil, nil
}

func UpdateTodoByID(id string, updatedTodo *models.Todo) (*models.Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i] = *updatedTodo
			return updatedTodo, nil
		}
	}
	return nil, nil
}

func DeleteTodoByID(id string) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return nil
}
