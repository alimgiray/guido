package user

import "github.com/alimgiray/guido/database"

type UserRepository struct {
	db *database.Database
}

func NewUserRepository(DB *database.Database) *UserRepository {
	return &UserRepository{
		db: DB,
	}
}
