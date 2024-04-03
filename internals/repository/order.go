// Order_repository.go
package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/muhaefath/to-do-list/internals/models"
)

type OrderRepository interface {
	GetAll() ([]models.Order, error)
	GetByID(id string) (*models.Order, error)
	CreateTodo(order *models.Order) error
	DeleteTodoByID(id string) error
}

// OrderRepository represents a repository for managing to-do list items
type orderRepository struct {
	db *sql.DB
}

// NewOrderRepository creates a new instance of OrderRepository
func NewOrderRepository(db *sql.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAll() ([]models.Order, error) {
	rows, err := r.db.Query("SELECT id, product_id, quantity FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Order []models.Order
	for rows.Next() {
		var item models.Order
		if err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity); err != nil {
			return nil, err
		}

		Order = append(Order, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Order, nil
}

func (r *orderRepository) GetByID(id string) (*models.Order, error) {
	rows, err := r.db.Query("SELECT id, product_id, quantity FROM orders WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var item models.Order
	for rows.Next() {
		if err := rows.Scan(&item.ID, &item.ProductID, &item.Quantity); err != nil {
			return nil, err
		}
	}

	return &item, nil
}

func (r *orderRepository) CreateTodo(order *models.Order) error {
	_, err := r.db.Exec("INSERT INTO orders (product_id, quantity) VALUES ($1, $2)", order.ProductID, order.Quantity)
	if err != nil {
		log.Println("Error creating orders:", err)
		return err
	}

	return nil
}

func (r *orderRepository) DeleteTodoByID(id string) error {
	_, err := r.db.Exec("DELETE FROM orders WHERE id = $1", id)
	if err != nil {
		log.Println("Error deleting orders:", err)
		return err
	}

	return nil
}
