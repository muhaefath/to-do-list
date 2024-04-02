package repository

import (
	"github.com/muhaefath/to-do-list/internals/models"
)

// MockTodoRepository represents a mock implementation of TodoListRepository for testing
type MockTodoRepository struct {
	GetAllFunc         func() ([]models.Todo, error)
	GetByIDFunc        func(id string) (*models.Todo, error)
	CreateTodoFunc     func(todo *models.Todo) error
	UpdateTodoByIDFunc func(id string, updatedTodo *models.Todo) error
	DeleteTodoByIDFunc func(id string) error
}

// GetAll retrieves all to-do items (mock implementation)
func (m *MockTodoRepository) GetAll() ([]models.Todo, error) {
	return m.GetAllFunc()
}

// GetByID retrieves a to-do item by ID (mock implementation)
func (m *MockTodoRepository) GetByID(id string) (*models.Todo, error) {
	return m.GetByIDFunc(id)
}

// CreateTodo creates a new to-do item (mock implementation)
func (m *MockTodoRepository) CreateTodo(todo *models.Todo) error {
	return m.CreateTodoFunc(todo)
}

// UpdateTodoByID updates a to-do item by ID (mock implementation)
func (m *MockTodoRepository) UpdateTodoByID(id string, updatedTodo *models.Todo) error {
	return m.UpdateTodoByIDFunc(id, updatedTodo)
}

// DeleteTodoByID deletes a to-do item by ID (mock implementation)
func (m *MockTodoRepository) DeleteTodoByID(id string) error {
	return m.DeleteTodoByIDFunc(id)
}
