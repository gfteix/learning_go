package address

import (
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

func (s *Store) CreateAddress(address types.Address, userID int) (int, error) {
	res, err := s.db.Exec("INSERT INTO addresses (user_id, street, city, address_state, postal_code, country) VALUES (?, ?, ?, ?, ?, ?)",
		userID, address.Street, address.City, address.State, address.PostalCode, address.Country)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (s *Store) GetAddress(addressId int) (*types.Address, error) {
	if addressId == 0 {
		return nil, fmt.Errorf("address ID cannot be 0")
	}

	row := s.db.QueryRow("SELECT id, user_id, street, city, address_state, postal_code, country FROM addresses WHERE id = ?", addressId)

	address := &types.Address{}
	err := row.Scan(&address.ID, &address.UserID, &address.Street, &address.City, &address.State, &address.PostalCode, &address.Country)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("address with ID %d not found", addressId)
		}
		return nil, err
	}

	return address, nil
}

func (s *Store) GetAddressesByUserID(userID int) ([]types.Address, error) {
	if userID == 0 {
		return nil, fmt.Errorf("user ID cannot be 0")
	}

	rows, err := s.db.Query("SELECT id, user_id, street, city, address_state, postal_code, country FROM addresses WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []types.Address

	for rows.Next() {
		address := types.Address{}
		err := rows.Scan(&address.ID, &address.UserID, &address.Street, &address.City, &address.State, &address.PostalCode, &address.Country)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return addresses, nil
}
