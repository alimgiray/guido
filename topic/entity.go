package topic

import "time"

type Topic struct {
	ID        int
	Title     string
	URL       string
	User      int
	Posts     []*Post
	CreatedAt time.Time
	UpdatedAt time.Time // Last post added time
}

type Post struct {
	ID        int
	Text      string
	User      int
	Topic     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
