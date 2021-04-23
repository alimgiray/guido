package main

import (
	"embed"
	"net/http"
	"path"

	"github.com/alimgiray/guido/config"
	"github.com/gin-gonic/gin"

	"github.com/alimgiray/guido/database"
	"github.com/alimgiray/guido/session"
	"github.com/alimgiray/guido/topic"
	"github.com/alimgiray/guido/user"
)

//go:embed ui/*
var f embed.FS

var userHandler *user.UserHandler
var topicHandler *topic.TopicHandler

func init() {
	db := database.Connect()

	userRepository := user.NewUserRepository(db)
	topicRepository := topic.NewTopicRepository(db)
	sessionRepository := session.NewSessionRepository(db)

	userService := user.NewUserService(userRepository)
	topicService := topic.NewTopicService(topicRepository)
	sessionService := session.NewSessionService(sessionRepository)

	configurationManager := config.NewConfigurationManager(db)

	userHandler = user.NewUserHandler(userService, sessionService, configurationManager)
	topicHandler = topic.NewTopicHandler(topicService, sessionService, configurationManager)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	if gin.Mode() == gin.ReleaseMode {
		r.HTMLRender = config.EmbedRenderer(f)
		r.GET("/css/*filepath", func(c *gin.Context) {
			c.FileFromFS(path.Join("/ui/assets/", c.Request.URL.Path), http.FS(f))
		})
		r.GET("/script/*filepath", func(c *gin.Context) {
			c.FileFromFS(path.Join("/ui/assets/", c.Request.URL.Path), http.FS(f))
		})
	} else {
		r.HTMLRender = config.LocalRenderer()
		r.Static("/css", "./ui/assets/css")
		r.Static("/script", "./ui/assets/script")
	}

	userRouter := r.Group("/user")
	userRouter.GET("/login", userHandler.GetLoginPage)
	userRouter.POST("/login", userHandler.Login)
	userRouter.GET("/register", userHandler.GetRegisterPage)
	userRouter.POST("/register", userHandler.Register)
	userRouter.GET("/logout", userHandler.Logout)

	topicRouter := r.Group("/topic")
	topicRouter.GET("/create", topicHandler.GetCreateTopicPage)
	topicRouter.POST("/create", topicHandler.CreateTopic)
	topicRouter.POST("/add", topicHandler.AddTopic)
	topicRouter.GET("/list", topicHandler.ListTopic)
	topicRouter.GET("/search", topicHandler.SearchTopic)

	r.GET("/:topic", topicHandler.GetTopic)
	r.GET("/", topicHandler.GetDefault)

	r.Run(":4000")
}
