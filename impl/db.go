package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConnection interface {
	Open(dbname string, host string, port int, pass string) error
	CreateUser(firstName string, lastName string, login string, passHash string, age int, phone int64) error
	AuthUser(login string, passHash string) (*User, error)
	GetFilms(pageSize int, pageNumber int, genre string, releaseYear int) (*[]Film, error)
	GetFilmsByUser(id int64, pageSize int, pageNumber int) (*[]Film, error)
	StartRent(login string, filmID int) error
	EndRent(filmID int, userID int64) error
	Close() error
}

type DB struct {
	connection *sql.DB
	connStrint string
}

func NewDB() *DB {
	return &DB{nil, ""}
}

func (db *DB) Open(dbname string, host string, port int, pass string) error {
	var err error
	connString := fmt.Sprintf("user=postgres dbname=%s host=%s port=%d password=%s sslmode=disable", dbname, host, port, pass)
	db.connection, err = sql.Open("postgres", connString)
	return err
}

func (db *DB) CreateUser(firstName string, lastName string, login string, passHash string, age int, phone int64) error {
	if db.connection == nil {
		return fmt.Errorf("Database connection not opened.")
	}

	_, err := db.connection.Query("SELECT public.create_user($1::VARCHAR, $2::VARCHAR, $3::VARCHAR, $4::VARCHAR, $5::INT, $6::BIGINT);", firstName, lastName, login, passHash, age, phone)
	return err
}
