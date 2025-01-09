package product

import (
	"database/sql"
	"ecommerce/types"
	"fmt"
	"strings"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT id, name, description, price, quantity, createdAt FROM products")

	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)

	for rows.Next() { // Advances to the next row in the result set
		p, err := scanRowIntoProduct(rows) // Processes the current row pointed to by 'rows'

		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetProductsByIDs(ids []int) ([]types.Product, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(ids)), ",")

	query := fmt.Sprintf("SELECT id, name, description, price, quantity, createdAt FROM products WHERE id IN (%s)", placeholders)

	args := make([]interface{}, len(ids))

	for i, v := range ids {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	products := []types.Product{}

	for rows.Next() {
		p, err := scanRowIntoProduct(rows)

		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = ?, price = ?, image = ?, description = ?, quantity = ? WHERE id = ?",
		product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID)

	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
