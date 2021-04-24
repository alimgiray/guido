package user

import (
	"errors"
	"time"

	"github.com/alimgiray/guido/database"
)

type UserRepository struct {
	db *database.Database
}

func NewUserRepository(DB *database.Database) *UserRepository {
	return &UserRepository{
		db: DB,
	}
}

// Insert a new user
func (u *UserRepository) Insert(username, password, email, userType string) error {
	query := "INSERT INTO user(username, email, password, createdAt, type) VALUES (?, ?, ?, ?, ?)"
	statement, err := u.db.Connection.Prepare(query)
	_, err = statement.Exec(username, email, password, time.Now().Format(time.RFC3339), userType)
	if err != nil {
		return err
	}
	return nil
}

// Find given user by username
func (u *UserRepository) FindByUsername(username string) (*User, error) {
	return u.Find("SELECT id, username, email, password, createdAt, type FROM user WHERE username = $1", username)
}

// Find given user by id
func (u *UserRepository) FindByID(userID int) (*User, error) {
	return u.Find("SELECT id, username, email, password, createdAt, type FROM user WHERE id = $1", userID)
}

func (u *UserRepository) Find(query string, args ...interface{}) (*User, error) {
	user := &User{}
	row := u.db.Connection.QueryRow(query, args[0])
	var createdAt string
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &createdAt, &user.Type)
	if err != nil {
		return nil, errors.New("user not found")
	}
	user.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	return user, nil
}
