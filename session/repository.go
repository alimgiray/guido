package session

import (
	"errors"
	"time"

	"github.com/alimgiray/guido/database"
)

type SessionRepository struct {
	db *database.Database
}

func NewSessionRepository(DB *database.Database) *SessionRepository {
	return &SessionRepository{
		db: DB,
	}
}

func (s *SessionRepository) FindUserIDBySessionID(sessionID string) (*Session, error) {
	session := &Session{}
	row := s.db.Connection.QueryRow("SELECT session_id, user, createdAt FROM session WHERE session_id = $1", sessionID)
	var createdAt string
	err := row.Scan(&session.SessionID, &session.UserID, &createdAt)
	if err != nil {
		return nil, errors.New("session not found")
	}
	session.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	return session, nil
}

func (s *SessionRepository) Insert(sessionID string, userID int) error {
	query := "INSERT INTO session(session_id, user, createdAt) VALUES (?, ?, ?)"
	statement, err := s.db.Connection.Prepare(query)
	_, err = statement.Exec(sessionID, userID, time.Now().Format(time.RFC3339))
	return err
}

func (s *SessionRepository) Delete(sessionID string) {
	query := "DELETE FROM session WHERE session_id = $1"
	statement, err := s.db.Connection.Prepare(query)
	if err == nil {
		statement.Exec(sessionID)
	}
}
