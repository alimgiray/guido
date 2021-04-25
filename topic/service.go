package topic

import (
	"log"
	"strings"
	"time"
)

type TopicService struct {
	topicRepository *TopicRepository
}

func NewTopicService(topicRepository *TopicRepository) *TopicService {
	return &TopicService{
		topicRepository: topicRepository,
	}
}

func (t *TopicService) CreateTopic(title, postText string, userID int) (string, error) {
	topic := &Topic{
		Title:     title,
		Posts:     make([]*Post, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		User:      userID,
		URL:       t.titleToURL(title),
	}
	topic, err := t.topicRepository.InsertTopic(topic)
	if err != nil {
		log.Println("TopicService-CreateTopic-InsertTopic", err.Error())
		return "", err
	}
	post := &Post{
		Text:      postText,
		Topic:     topic.ID,
		CreatedAt: time.Now(),
		User:      userID,
	}
	post, err = t.topicRepository.InsertPost(post)
	if err != nil {
		// TODO also delete the topic
		log.Println("TopicService-CreateTopic-InsertPost", err.Error())
		return "", err
	}
	topic.Posts = append(topic.Posts, post)
	return topic.URL, nil
}

func (s *TopicService) titleToURL(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}
