package topic

import (
	"github.com/gin-gonic/gin"

	"github.com/alimgiray/guido/config"
	"github.com/alimgiray/guido/session"
)

type TopicHandler struct {
	topicService   *TopicService
	sessionService *session.SessionService
	config         *config.ConfigurationManager
}

func NewTopicHandler(
	topicService *TopicService,
	sessionService *session.SessionService,
	configurationManager *config.ConfigurationManager) *TopicHandler {
	return &TopicHandler{
		topicService:   topicService,
		sessionService: sessionService,
		config:         configurationManager,
	}
}

func (t *TopicHandler) GetCreateTopicPage(c *gin.Context) {
	c.HTML(200, "create", gin.H{
		"Title": "Create New Topic",
	})
}

func (t *TopicHandler) CreateTopic(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func (t *TopicHandler) AddTopic(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func (t *TopicHandler) ListTopic(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func (t *TopicHandler) GetTopic(c *gin.Context) {
	c.HTML(200, "topic", gin.H{
		"Title": "Some topic",
	})
}

func (t *TopicHandler) GetDefault(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func (t *TopicHandler) SearchTopic(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
