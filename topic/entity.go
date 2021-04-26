package topic

import "time"

type Topic struct {
	ID        int
	Title     string
	URL       string
	User      int
	Username  string
	Posts     []*Post
	CreatedAt time.Time
	UpdatedAt time.Time // Last post added time
}

type Post struct {
	ID        int
	Text      string
	User      int
	Username  string
	Topic     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Post) FormatDate(t time.Time) string {
	return t.Format("2006-01-02 15:04")
}

func (p *Post) IsUpdated() bool {
	var defaultTime time.Time
	return p.UpdatedAt != defaultTime
}
