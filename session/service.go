package session

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

type SessionService struct {
	sessionRepository *SessionRepository
}

func NewSessionService(sessionRepository *SessionRepository) *SessionService {
	return &SessionService{
		sessionRepository: sessionRepository,
	}
}

// IsUserLoggedIn returns userID
func (s *SessionService) IsUserLoggedIn(sessionID string) (int, bool) {
	session, err := s.sessionRepository.FindUserIDBySessionID(sessionID)
	if err != nil || session.UserID == 0 {
		return 0, false
	}
	return session.UserID, true
}

// CreateSession returns sessionID
func (s *SessionService) CreateSession(userID int) (string, error) {
	sessionID, err := s.createSessionID()
	if err != nil {
		return "", err
	}
	err = s.sessionRepository.Insert(sessionID, userID)
	if err != nil {
		return "", err
	}
	return sessionID, nil
}

// RemoveSession deletes sessionID from database
func (s *SessionService) RemoveSession(sessionID string) {
	s.sessionRepository.Delete(sessionID)
}

func (s *SessionService) createSessionID() (string, error) {
	sessionID := make([]byte, 128)
	_, err := rand.Read(sessionID)
	if err != nil {
		log.Println("SessionService-createSessionID", err)
		return "", err
	}
	return base64.URLEncoding.EncodeToString(sessionID), nil
}
