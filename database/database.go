package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import the driver
)

var name = "guido.db"

type Database struct {
	Connection *sql.DB
}

func Connect() *Database {
	var db *sql.DB
	if !(exists(name)) {
		db = create()
	} else {
		db = open()
	}
	return &Database{
		Connection: db,
	}
}

func create() *sql.DB {
	log.Println("First time run, creating a new database")
	_, err := os.Create(name)
	checkError(err)
	db := open()

	createTable(db, createUserTableQuery)
	createTable(db, createTopicTableQuery)
	createTable(db, createPostTableQuery)
	createTable(db, createSessionTableQuery)
	createTable(db, createConfigTableQuery)

	createAdmin(db)
	createConfig(db)
	log.Println("Database created successfully")

	return db
}

func createTable(db *sql.DB, query string) {
	statement, err := db.Prepare(query)
	checkError(err)
	_, err = statement.Exec()
	checkError(err)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func open() *sql.DB {
	db, err := sql.Open("sqlite3", name)
	checkError(err)
	return db
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
