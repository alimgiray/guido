package topic

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/alimgiray/guido/database"
)

type TopicRepository struct {
	db *database.Database
}

func NewTopicRepository(DB *database.Database) *TopicRepository {
	return &TopicRepository{
		db: DB,
	}
}

func (t *TopicRepository) FindByURL(url string) (*Topic, error) {
	topic := &Topic{}
	row := t.db.Connection.QueryRow("SELECT id, title, url, user, createdAt, updatedAt FROM topic WHERE url = $1", url)
	var createdAt string
	var updatedAt sql.NullString
	err := row.Scan(&topic.ID, &topic.Title, &topic.URL, &topic.User, &createdAt, &updatedAt)
	if err != nil {
		log.Println("TopicRepository-FindByURL", err.Error())
		return nil, errors.New("not found")
	}
	topic.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
	if updatedAt.Valid {
		topic.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt.String)
	}
	return topic, nil
}

func (t *TopicRepository) InsertTopic(topic *Topic) (*Topic, error) {
	query := "INSERT INTO topic(title, url, user, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?)"
	statement, err := t.db.Connection.Prepare(query)
	if err != nil {
		log.Println("TopicRepository-InsertTopic-Prepare", err.Error())
		return nil, err
	}
	result, err := statement.Exec(&topic.Title, &topic.URL, &topic.User, &topic.CreatedAt, &topic.UpdatedAt)
	if err != nil {
		log.Println("TopicRepository-InsertTopic-Exec", err.Error())
		return nil, err
	}
	topicID, _ := result.LastInsertId()
	topic.ID = int(topicID)
	return topic, nil
}

func (t *TopicRepository) UpdateTopic(topic *Topic) error {
	query := "UPDATE topic SET title = ?, updatedAt = ? WHERE id = ?"
	statement, err := t.db.Connection.Prepare(query)
	if err != nil {
		log.Println("TopicRepository-UpdateTopic-Prepare", err.Error())
		return err
	}
	_, err = statement.Exec(&topic.Title, &topic.UpdatedAt, &topic.ID)
	if err != nil {
		log.Println("TopicRepository-UpdateTopic-Exec", err.Error())
		return err
	}
	return nil
}

func (t *TopicRepository) GetTopics(criteria, order string) []*Topic {
	topics := make([]*Topic, 0, 20)
	rows, err := t.db.Connection.Query(fmt.Sprintf("SELECT title, url FROM topic ORDER BY %s %s", criteria, order))
	if err != nil {
		log.Println("TopicRepository-GetTopics-Query", err.Error())
		return topics
	}
	for true {
		if !rows.Next() {
			break
		}
		topic := &Topic{}
		rows.Scan(&topic.Title, &topic.URL)
		topics = append(topics, topic)
	}
	return topics
}

func (t *TopicRepository) InsertPost(post *Post) (*Post, error) {
	query := "INSERT INTO post(text, user, topic, createdAt) VALUES (?, ?, ?, ?)"
	statement, err := t.db.Connection.Prepare(query)
	if err != nil {
		log.Println("TopicRepository-InsertTopic", err.Error())
		return nil, err
	}
	result, err := statement.Exec(&post.Text, &post.User, &post.Topic, &post.CreatedAt)
	if err != nil {
		log.Println("TopicRepository-InsertPost", err.Error())
		return nil, err
	}
	postID, _ := result.LastInsertId()
	post.ID = int(postID)
	return post, nil
}

// TODO pagination
func (t *TopicRepository) FindPostsByTopic(topicID int) ([]*Post, error) {
	posts := make([]*Post, 0, 20)
	rows, err := t.db.Connection.Query("SELECT id, text, user, createdAt, updatedAt FROM post WHERE topic = $1", topicID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for true {
		if !rows.Next() {
			break
		}
		post := &Post{}
		var createdAt string
		var updatedAt sql.NullString
		rows.Scan(&post.ID, &post.Text, &post.User, &createdAt, &updatedAt)
		post.CreatedAt, _ = time.Parse(time.RFC3339, createdAt)
		if updatedAt.Valid {
			post.UpdatedAt, _ = time.Parse(time.RFC3339, updatedAt.String)
		}
		posts = append(posts, post)
	}
	return posts, nil
}
