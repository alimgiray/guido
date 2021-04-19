package session

import "time"

type Session struct {
	SessionID string
	UserID    int
	CreatedAt time.Time
}
