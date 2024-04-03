package product

import (
	"database/sql"
	"log"

	"github.com/muhaefath/to-do-list/internals/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id string) (*models.Product, error)
	UpdateQuantityByID(id string, product *models.Product) error
}

type productRepository struct {
	db *sql.DB
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	rows, err := r.db.Query("SELECT id, title, quantity FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Product []models.Product
	for rows.Next() {
		var item models.Product
		if err := rows.Scan(&item.ID, &item.Title, &item.Quantity); err != nil {
			return nil, err
		}

		Product = append(Product, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return Product, nil
}

func (r *productRepository) GetByID(id string) (*models.Product, error) {
	rows, err := r.db.Query("SELECT id, title, quantity FROM products WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var item models.Product
	for rows.Next() {
		if err := rows.Scan(&item.ID, &item.Title, &item.Quantity); err != nil {
			return nil, err
		}
	}

	return &item, nil
}

func (r *productRepository) UpdateQuantityByID(id string, product *models.Product) error {
	_, err := r.db.Exec("UPDATE products SET quantity = $1 WHERE id = $2", product.Quantity, id)
	if err != nil {
		log.Println("Error updating todo:", err)
		return err
	}

	return nil
}
