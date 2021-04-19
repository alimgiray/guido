package session

type SessionService struct {
	sessionRepository *SessionRepository
}

func NewSessionService(sessionRepository *SessionRepository) *SessionService {
	return &SessionService{
		sessionRepository: sessionRepository,
	}
}
