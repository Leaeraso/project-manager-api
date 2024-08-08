package main

import "database/sql"

type Store interface {
	// Users
	CreateUser() error
	// Tasks
	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
}

// Tendra los metodos para comunicarse con la base de datos
type Storage struct {
	db *sql.DB
}

// Constructor
func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}
}

// Method
func (s *Storage) CreateUser() error {
	return nil
}

func (s *Storage) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec(`insert into tasks (name, status, project_id, assigned_to) values 
	(?, ?, ?, ?)`, t.Name, t.Status, t.ProjectID, t.AssignedToID)

	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}

func (s *Storage) GetTask(id string) (*Task, error) {
	var t Task
	err := s.db.QueryRow("select id, name, status, project_id, assigned_to, createdAt from tasks where id = ?", id).Scan(&t.ID, &t.Name, &t.Status, &t.ProjectID, &t.AssignedToID, &t.CreatedAt)

	return &t, err
}
