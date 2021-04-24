package database

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

const createUserTableQuery = `
	CREATE TABLE user (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
		"username" TEXT NOT NULL,
		"email" TEXT,
		"password" TEXT NOT NULL,
		"createdAt" TEXT NOT NULL,
		"type" TEXT NOT NULL
	);
`

func createAdmin(db *sql.DB) {
	fmt.Println("Creating admin user")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		os.Remove(name)
		log.Fatal(err)
	}
	username = strings.Trim(username, "\n")

	fmt.Print("Enter Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		os.Remove(name)
		log.Fatal(err)
	}
	email = strings.Trim(username, "\n")

	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		os.Remove(name)
		log.Fatal(err)
	}
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	password := string(passwordHash)

	query := `INSERT INTO user(username, email, password, createdAt, type) VALUES (?, ?, ?, ?, ?)`
	statement, err := db.Prepare(query)
	_, err = statement.Exec(username, email, password, time.Now().Format(time.RFC3339), "admin")
	if err != nil {
		log.Fatal(err.Error())
	}
}
