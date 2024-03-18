package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func Connect() {
	// Create or open the SQLite database
	var err error
	DB, err = sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create a table if it doesn't exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS categories (
    category TEXT PRIMARY KEY NOT NULL
)`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS memos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    category TEXT NOT NULL,
    date TEXT NOT NULL,
    FOREIGN KEY (category) REFERENCES categories(category)
)`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
