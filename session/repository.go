package session

import "github.com/alimgiray/guido/database"

type SessionRepository struct {
	db *database.Database
}

func NewSessionRepository(DB *database.Database) *SessionRepository {
	return &SessionRepository{
		db: DB,
	}
}
