package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

//Esta es la implementacion para MySQL
type MySQLStorage struct {
	db *sql.DB
}

//Constructor
func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL")

	return &MySQLStorage{db: db}
}

//Method
func (ms *MySQLStorage) Init() (*sql.DB, error) {
	//En este metodo inicializamos las tablas de la base de datos y luego devolvemos un puntero a esa base de datos.
	return ms.db, nil
}
