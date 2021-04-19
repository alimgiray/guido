package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Import the driver
)

var name = "guido.db"

type Database struct {
	Connection *sql.DB
}

func Connect() *Database {
	return nil
}
