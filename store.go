package main

import "database/sql"

type Store interface {
	// Users
	CreateUser() error
}

// Tendra los metodos para comunicarse con la base de datos
type Storage struct {
	db *sql.DB
}
//Constructor
func NewStore(db *sql.DB) *Storage {
	return &Storage{db: db}
}
//Method
func (s *Storage) CreateUser() error {
	return nil
}
