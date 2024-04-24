package main

import "database/sql"

type Store interface {
	GetUserById(id string) (*User, error)
	CreateUser(u *User) (*User, error)
	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
	CreateProject(p *Project) (*Project, error)
	GetProject(id string) (*Project, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) GetUserById(id string) (*User, error) {
	var u User
	err := s.db.QueryRow("SELECT id, email, name, firstName, lastName, createdAt FROM users WHERE id = ?", id).Scan(&u.Id, &u.Email, &u.Name, &u.FirstName, &u.LastName, &u.CreatedAt)
	return &u, err
}

func (s *Storage) CreateUser(u *User) (*User, error) {
	rows, err := s.db.Exec("INSERT INTO users (email, name, firstName, lastName, password) VALUES (?, ?, ?, ?, ?)",
		u.Email, u.Name, u.FirstName, u.LastName, u.Password)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()

	if err != nil {
		return nil, err
	}

	u.Id = id
	return u, nil
}

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec("INSERT INTO tasks (name, status, project_id, assigned_to) VALUES (?, ?, ?, ?)",
		t.Name, t.Status, t.ProjectId, t.AssignedToId)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()

	if err != nil {
		return nil, err
	}

	t.Id = id
	return t, nil
}

func (s *Storage) GetTask(id string) (*Task, error) {
	var t Task
	rows := s.db.QueryRow("SELECT id, name, status, project_id, assigned_to, createdAt FROM tasks WHERE id = ?", id)
	err := rows.Scan(&t.Id, &t.Name, &t.Status, &t.ProjectId, &t.AssignedToId, &t.CreatedAt)
	return &t, err
}

func (s *Storage) CreateProject(p *Project) (*Project, error) {
	rows, err := s.db.Exec("INSERT INTO projects (name) VALUES (?)",
		p.Name,
	)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()

	if err != nil {
		return nil, err
	}

	p.Id = id
	return p, nil
}

func (s *Storage) GetProject(id string) (*Project, error) {
	var p Project
	rows := s.db.QueryRow("SELECT id, name, createdAt FROM projects WHERE id = ?", id)
	err := rows.Scan(&p.Id, &p.Name, &p.CreatedAt)
	return &p, err
}
