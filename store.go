package main

import "database/sql"

type Store interface {
	// Users
	CreateUser(u *User) (*User, error)
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	// Tasks
	CreateTask(t *Task) (*Task, error)
	GetTask(id string) (*Task, error)
	// Projects
	CreateProject(p *Project) (*Project, error)
	GetProjectById(id string) (*Project, error)
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
func (s *Storage) CreateUser(u *User) (*User, error) {
	rows, err := s.db.Exec(`insert into users (email, firstName, lastName, password) values 
	(?, ?, ?, ?)`, u.Email, u.FirstName, u.LastName, u.Password)
	
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	u.ID = id
	return u, nil
}

func (s *Storage) GetUserByID(id string) (*User, error) {
	var u User
	err := s.db.QueryRow("select id, firstName, lastName, email, createdAt from users where id = ?", id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.CreatedAt)

	return &u, err
}

func (s *Storage) GetUserByEmail(email string) (*User, error) {
	var u User
	err := s.db.QueryRow("select id, email, password from users where email = ?", email).Scan(&u.ID, &u.Email, &u.Password)

	return &u, err
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

func (s *Storage) CreateProject(p *Project) (*Project, error) {
	rows, err := s.db.Exec(`insert into projects (name) values 
	(?)`, p.Name)
	
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	p.ID = id
	return p, nil
}

func (s *Storage) GetProjectById(id string) (*Project, error) {
	var p Project
	err := s.db.QueryRow("select id, name, createdAt from projects where id = ?", id).Scan(&p.ID, &p.Name, &p.CreatedAt)

	return &p, err
}
