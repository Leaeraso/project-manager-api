package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Esta es la implementacion para MySQL
type MySQLStorage struct {
	db *sql.DB
}

var db *sql.DB
var err error

// Constructor
func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	for {
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Printf("Error opening the database: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("cannot connect to the database: %v", err)
			db.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		break
	}

	log.Println("Connected to MySQL")

	return &MySQLStorage{db: db}
}

// Method
func (ms *MySQLStorage) Init() (*sql.DB, error) {
	//En este metodo inicializamos las tablas de la base de datos y luego devolvemos un puntero a esa base de datos.

	if err := ms.createProjectsTable(); err != nil {
		return nil, err
	}

	if err := ms.createUsersTable(); err != nil {
		return nil, err
	}

	if err := ms.createTasksTable(); err != nil {
		return nil, err
	}

	return ms.db, nil
}

func (ms *MySQLStorage) createProjectsTable() error {
	_, err := ms.db.Exec(`
		create table is not exists projects (
			id int unsigned not null auto_increment,
			name varchar(255) not null,
			createdAt timestamp not null default current_timestamp,

			primary key(id),
		) engine=InnoDB default charset=utf8;
	`)

	return err
}

func (ms *MySQLStorage) createTasksTable() error {
	_, err := ms.db.Exec(`
		create table if not exists tasks (
			id int unsigned not null auto_increment,
			name varchar(255) not null,
			status enum('TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE') not null default 'TODO',
			projectId int unsigned not null,
			assignedToID int unsigned not null,
			createdAt timestamp not null default current_timestamp,

			primary key(id),
			foreign key(assignedToID) references users(id),
			foreign key(projectId) references projects(id)
		) engine=InnoDB default charset=utf8;
	`)

	if err != nil {
		return err
	}

	return nil
}

func (ms *MySQLStorage) createUsersTable() error {
	_, err := ms.db.Exec(`
		create table if not exists users (
			id int unsigned not null auto_increment,
			email varchar(255) not null,
			firstName varchar(255) not null,
			lastName varchar(255) not null,
			password varchar(255) not null,
			createdAt timestamp not null default current_timestamp,

			primary key(id),
			unique key(email),
		) engine=InnoDB default charset=utf8;
	`)

	return err
}
