package topic

import "github.com/alimgiray/guido/database"

type TopicRepository struct {
	db *database.Database
}

func NewTopicRepository(DB *database.Database) *TopicRepository {
	return &TopicRepository{
		db: DB,
	}
}
