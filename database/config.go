package database

import (
	"database/sql"
	"fmt"
	"log"
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
	defaultURL(db)
}

func appName(db *sql.DB) {
	appName := "Guido"
	query := `INSERT INTO config(name, value, createdAt, updatedAt) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(query)
	_, err = statement.Exec("appName", appName, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	if err != nil {
		log.Println(err)
	}
}

func userType(db *sql.DB) {
	userType := "user"
	query := `INSERT INTO config(name, value, createdAt, updatedAt) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(query)
	_, err = statement.Exec("userType", userType, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	if err != nil {
		log.Println(err)
	}
}

func defaultURL(db *sql.DB) {
	defaultURL := "guido"
	query := `INSERT INTO config(name, value, createdAt, updatedAt) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(query)
	_, err = statement.Exec("defaultURL", defaultURL, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
	if err != nil {
		log.Println(err)
	}
}
