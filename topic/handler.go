package topic

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/alimgiray/guido/user"

	"github.com/gin-gonic/gin"

	"github.com/alimgiray/guido/config"
	"github.com/alimgiray/guido/session"
)

type TopicHandler struct {
	topicService   *TopicService
	sessionService *session.SessionService
	userService    *user.UserService
	config         *config.ConfigurationManager
}

func NewTopicHandler(
	topicService *TopicService,
	sessionService *session.SessionService,
	userService *user.UserService,
	configurationManager *config.ConfigurationManager) *TopicHandler {
	return &TopicHandler{
		topicService:   topicService,
		sessionService: sessionService,
		userService:    userService,
		config:         configurationManager,
	}
}

func (t *TopicHandler) GetCreateTopicPage(c *gin.Context) {
	username, err := t.getUsernameFromCookie(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", t.config.GetDefaultURL()))
		return
	}
	c.HTML(200, "create", gin.H{
		"Title":  "Create New Topic",
		"Header": t.config.GetHeader(username, true),
	})
}

type CreateTopicForm struct {
	Title string `form:"title" binding:"required"`
	Post  string `form:"post" binding:"required"`
}

func (t *TopicHandler) CreateTopic(c *gin.Context) {
	userID, err := t.getUserID(c)
	if err != nil {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", t.config.GetDefaultURL()))
		return
	}

	var form CreateTopicForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO error page
		return
	}

	url, err := t.topicService.CreateTopic(form.Title, form.Post, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO error page
		return
	}

	// Topic created, redirect user to that topic
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", url))
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
	log.Println("get topic")
	log.Println(c.Param("topic"))
	// Get default header
	header := t.config.GetHeader("", false)
	username, err := t.getUsernameFromCookie(c)
	if err == nil {
		header = t.config.GetHeader(username, true)
	}
	c.HTML(200, "topic", gin.H{
		"Title":       "Some topic",
		"Meta":        t.config.GetMeta("desc", "k1, k2, k3"),
		"Header":      header,
		"ShowSidebar": true,
	})
}

func (t *TopicHandler) GetDefault(c *gin.Context) {
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", t.config.GetDefaultURL()))
}

func (t *TopicHandler) SearchTopic(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func (t *TopicHandler) getUserID(c *gin.Context) (int, error) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return 0, err
	}
	userID, loggedIn := t.sessionService.IsUserLoggedIn(cookie)
	if loggedIn {
		return userID, nil
	}
	return 0, errors.New("user not found")
}

func (t *TopicHandler) getUsernameFromCookie(c *gin.Context) (string, error) {
	userID, err := t.getUserID(c)
	if err != nil {
		return "", err
	}

	return t.userService.GetUsernameFromID(userID)
}
