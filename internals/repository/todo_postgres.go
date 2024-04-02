// todolist_repository.go
package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/muhaefath/to-do-list/internals/models"
)

// TodoListRepository represents a repository for managing to-do list items
type TodoListRepository struct {
	db *sql.DB
}

// NewTodoListRepository creates a new instance of TodoListRepository
func NewTodoListRepository(db *sql.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

// GetAll retrieves all to-do items from the database
func (r *TodoListRepository) GetAll() ([]models.Todo, error) {
	rows, err := r.db.Query("SELECT id, title, completed FROM todolist")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todoList []models.Todo
	for rows.Next() {
		var item models.Todo
		if err := rows.Scan(&item.ID, &item.Title, &item.Completed); err != nil {
			return nil, err
		}

		todoList = append(todoList, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todoList, nil
}

// GetAll retrieves all to-do items from the database
func (r *TodoListRepository) GetByID(id string) (*models.Todo, error) {
	rows, err := r.db.Query("SELECT id, title, completed FROM todolist WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var item models.Todo
	for rows.Next() {
		if err := rows.Scan(&item.ID, &item.Title, &item.Completed); err != nil {
			return nil, err
		}
	}

	return &item, nil
}

func (r *TodoListRepository) CreateTodo(todo *models.Todo) error {
	_, err := r.db.Exec("INSERT INTO todolist (title, completed) VALUES ($1, $2)", todo.Title, todo.Completed)
	if err != nil {
		log.Println("Error creating todo:", err)
		return err
	}

	return nil
}

func (r *TodoListRepository) UpdateTodoByID(id string, updatedTodo *models.Todo) error {
	_, err := r.db.Exec("UPDATE todolist SET title = $1, completed = $2 WHERE id = $3", updatedTodo.Title, updatedTodo.Completed, id)
	if err != nil {
		log.Println("Error updating todo:", err)
		return err
	}

	return nil
}

func (r *TodoListRepository) DeleteTodoByID(id string) error {
	_, err := r.db.Exec("DELETE FROM todolist WHERE id = $1", id)
	if err != nil {
		log.Println("Error deleting todo:", err)
		return err
	}

	return nil
}
