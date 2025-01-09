package order

import (
	"context"
	"database/sql"
	"ecommerce/types"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(ctx context.Context, productMap map[int]types.Product, cart types.CartCheckoutPayload, userID int) (int, float64, error) {
	fail := func(tx *sql.Tx, err error) (int, float64, error) {
		tx.Rollback()
		return 0, 0, fmt.Errorf("CreateOrder: %v", err)
	}

	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		return fail(tx, err)
	}

	total := calculateTotalPrice(cart.Items, productMap)

	var addressId int

	if cart.AddressId != nil {
		addressId = *cart.AddressId
	} else {
		res, err := tx.ExecContext(ctx, "INSERT INTO addresses (user_id, street, city, address_state, postal_code, country) VALUES (?, ?, ?, ?, ?, ?)",
			userID, cart.Address.Street, cart.Address.City, cart.Address.State, cart.Address.PostalCode, cart.Address.Country)

		if err != nil {
			return fail(tx, err)
		}

		id, err := res.LastInsertId()

		if err != nil {
			return fail(tx, err)
		}

		addressId = int(id)
	}

	if addressId == 0 {
		return fail(tx, err)
	}

	res, err := tx.ExecContext(ctx, "INSERT INTO orders (user_id, total, status, address_id) VALUES (?, ?, ?, ?)", userID, total, "pending", addressId)

	if err != nil {
		return fail(tx, err)
	}

	orderId, err := res.LastInsertId()

	if err != nil {
		return fail(tx, err)
	}

	for _, cartItem := range cart.Items {
		p := productMap[cartItem.ProductID]
		qty := p.Quantity - cartItem.Quantity

		_, err = tx.ExecContext(ctx, "INSERT INTO order_items (order_id, product_id, quantity, price) VALUES(?, ?, ?, ?)",
			int(orderId), cartItem.ProductID, cartItem.Quantity, p.Price)

		if err != nil {
			return fail(tx, err)
		}

		_, err := tx.ExecContext(ctx, "UPDATE products SET quantity = ? WHERE id = ?;", qty, cartItem.ProductID)

		if err != nil {
			return fail(tx, err)
		}

	}

	if err = tx.Commit(); err != nil {
		return fail(tx, err)
	}

	return int(orderId), total, nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES(?, ?, ?, ?)",
		orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)

	return err
}

func (s *Store) GetOrdersByUserId(userID int) ([]types.Order, error) {
	if userID == 0 {
		return nil, nil
	}

	rows, err := s.db.Query("SELECT id, user_id, total, status, address_id, created_at FROM orders WHERE user_id = ? ", userID)

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

func (s *Store) UpdateOrder(orderID int, status string) error {
	_, err := s.db.Exec("UPDATE orders SET status = ? WHERE id = ?;`", status, orderID)

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

func calculateTotalPrice(cartItems []types.CartItemPayload, products map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
