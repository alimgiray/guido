package user

import "time"

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	Type      string
}
