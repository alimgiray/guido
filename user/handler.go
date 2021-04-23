package user

import (
	"github.com/alimgiray/guido/config"
	"github.com/alimgiray/guido/session"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService    *UserService
	sessionService *session.SessionService
	config         *config.ConfigurationManager
}

func NewUserHandler(
	userService *UserService,
	sessionService *session.SessionService,
	configurationManager *config.ConfigurationManager) *UserHandler {
	return &UserHandler{
		userService:    userService,
		sessionService: sessionService,
		config:         configurationManager,
	}
}

func (u *UserHandler) GetLoginPage(c *gin.Context) {
	c.HTML(200, "login", gin.H{
		"Title":  "Login",
		"Header": u.config.GetHeader("", false),
	})
}

func (u *UserHandler) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func (u *UserHandler) GetRegisterPage(c *gin.Context) {
	c.HTML(200, "register", gin.H{
		"Title":  "Register",
		"Header": u.config.GetHeader("", false),
	})
}

func (u *UserHandler) Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func (u *UserHandler) Logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
