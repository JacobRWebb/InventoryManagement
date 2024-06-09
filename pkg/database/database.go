package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	dbFileName := "database.db"

	if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
		file, err := os.Create(dbFileName)

		if err != nil {
			log.Fatalf("Unable to create database file. %v", err)
		}

		file.Close()
	}

	var err error
	DB, err = sql.Open("sqlite3", dbFileName)

	if err != nil {
		log.Fatalf("Error opening database. %v", err)
	}

	seedTables()
}

func seedTables() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
            id TEXT PRIMARY KEY,
            email TEXT NOT NULL UNIQUE,
            passwordHash TEXT NOT NULL
        )
	`)

	if err != nil {
		log.Fatalf("Unable to execute, %v", err)
	}
}
