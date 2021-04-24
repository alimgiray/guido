package database

import (
	"database/sql"
	"fmt"
	"time"
)

const createConfigTableQuery = `
	CREATE TABLE config (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		"name" TEXT NOT NULL,
		"value" TEXT NOT NULL,
		"createdAt" TEXT NOT NULL,
		"updatedAt" TEXT
	);
`

func createConfig(db *sql.DB) {
	fmt.Println("Creating configurations...")
	appName(db)
	userType(db)
}

func appName(db *sql.DB) error {
	appName := "Guido"
	query := `INSERT INTO config(name, value, createdAt, updatedAt) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(query)
	_, err = statement.Exec("appName", appName, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	return err
}

func userType(db *sql.DB) error {
	userType := "user"
	query := `INSERT INTO config(name, value, createdAt, updatedAt) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(query)
	_, err = statement.Exec("userType", userType, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	return err
}
