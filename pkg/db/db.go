package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase() *Database {
	conn, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database %v", err)
	}

	query := `
	CREATE TABLE IF NOT exists users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
	);`

	if _, err := conn.Exec(query); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	return &Database{Conn: conn}
}

