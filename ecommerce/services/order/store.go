package order

import (
	"context"
	"database/sql"
	"ecommerce/types"
	"errors"
	"fmt"
	"log"
	"strings"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(ctx context.Context, productIDs []int, cart types.CartCheckoutPayload, userID int) (int, float64, error) {
	fail := func(tx *sql.Tx, err error) (int, float64, error) {
		_ = tx.Rollback()
		return 0, 0, fmt.Errorf("CreateOrder: %v", err)
	}

	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		return fail(nil, err)
	}

	addressID, err := getOrderAddressId(tx, ctx, cart, userID)

	if err != nil {
		return fail(tx, err)
	}

	productMap, err := getProductsToUpdate(tx, ctx, productIDs)

	if err != nil {
		return fail(nil, err)
	}

	total := calculateTotalPrice(cart.Items, productMap)

	res, err := tx.ExecContext(ctx, "INSERT INTO orders (user_id, total, status, address_id) VALUES (?, ?, ?, ?)", userID, total, "pending", addressID)

	if err != nil {
		return fail(tx, err)
	}

	orderID, err := res.LastInsertId()

	if err != nil {
		return fail(tx, err)
	}

	err = createOrderItems(tx, ctx, cart, productMap, int(orderID))

	if err != nil {
		return fail(tx, err)
	}

	if err = tx.Commit(); err != nil {
		return fail(tx, err)
	}

	return int(orderID), total, nil
}

func (s *Store) GetOrdersByUserId(userID int) ([]types.Order, error) {
	if userID == 0 {
		return nil, nil
	}

	rows, err := s.db.Query("SELECT id, user_id, total, status, address_id, created_at FROM orders WHERE user_id = ?", userID)

	if err != nil {
		return nil, err
	}

	orders := []types.Order{}

	for rows.Next() {
		o, err := scanRowIntoOrder(rows)

		if err != nil {
			return nil, err
		}

		orders = append(orders, *o)
	}

	return orders, nil
}

func (s *Store) GetOrder(id int) (*types.Order, error) {
	rows, err := s.db.Query("SELECT id, user_id, total, status, address_id, created_at FROM orders WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	o := new(types.Order)

	for rows.Next() {
		o, err = scanRowIntoOrder(rows)

		if err != nil {
			return nil, err
		}
	}

	if o.ID == 0 {
		return nil, nil
	}

	return o, nil
}

func (s *Store) UpdateOrder(orderID int, status string) error {
	log.Printf("%v %v", orderID, status)
	_, err := s.db.Exec("UPDATE orders SET status = ? WHERE id = ?;", status, orderID)

	if err != nil {
		return err
	}

	return nil
}

func scanRowIntoOrder(rows *sql.Rows) (*types.Order, error) {
	order := new(types.Order)

	err := rows.Scan(
		&order.ID,
		&order.UserID,
		&order.Total,
		&order.Status,
		&order.AddressId,
		&order.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func getOrderAddressId(tx *sql.Tx, ctx context.Context, cart types.CartCheckoutPayload, userID int) (int, error) {
	if cart.AddressId != nil {
		return *cart.AddressId, nil
	}

	if cart.Address != nil {
		res, err := tx.ExecContext(ctx, "INSERT INTO addresses (user_id, street, city, address_state, postal_code, country) VALUES (?, ?, ?, ?, ?, ?)",
			userID, cart.Address.Street, cart.Address.City, cart.Address.State, cart.Address.PostalCode, cart.Address.Country)
		if err != nil {
			return 0, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return 0, err
		}

		return int(id), nil
	}

	return 0, errors.New("no address")
}

func getProductsToUpdate(tx *sql.Tx, ctx context.Context, productIDs []int) (map[int]types.Product, error) {
	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(productIDs)), ",")

	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
	}

	query := fmt.Sprintf("SELECT id, name, description, price, quantity, created_at FROM products WHERE id IN (%s) FOR UPDATE", placeholders)

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []types.Product{}

	for rows.Next() {
		p := new(types.Product)
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Quantity,
			&p.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	productMap := make(map[int]types.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	return productMap, nil
}

func createOrderItems(tx *sql.Tx, ctx context.Context, cart types.CartCheckoutPayload, productMap map[int]types.Product, orderID int) error {
	for _, cartItem := range cart.Items {
		product := productMap[cartItem.ProductID]
		newQuantity := product.Quantity - cartItem.Quantity

		_, err := tx.ExecContext(ctx, "INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (?, ?, ?, ?)",
			orderID, cartItem.ProductID, cartItem.Quantity, product.Price)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, "UPDATE products SET quantity = ? WHERE id = ?", newQuantity, cartItem.ProductID)
		if err != nil {
			return err
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []types.CartItemPayload, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
