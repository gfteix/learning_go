package user

import (
	"database/sql"
	"ecommerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan( // Scans data from the current row
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT id, first_name, last_name, email, password, created_at FROM users WHERE email = ?", email)

	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() { // Advances to the next row in the result set
		u, err = scanRowIntoUser(rows) // Processes the current row pointed to by 'rows'

		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, nil // if no user, returns nil
	}

	return u, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() { // Advances to the next row in the result set
		u, err = scanRowIntoUser(rows) // Processes the current row pointed to by 'rows'

		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, nil
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)",
		user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}
