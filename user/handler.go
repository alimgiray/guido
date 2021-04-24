package user

import (
	"fmt"
	"net/http"

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
	if u.isLoggedIn(c) {

		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", u.config.GetDefaultURL()))
		return
	}

	c.HTML(200, "login", gin.H{
		"Title":  "Login",
		"Header": u.config.GetHeader("", false),
	})
}

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (u *UserHandler) Login(c *gin.Context) {
	if u.isLoggedIn(c) {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", u.config.GetDefaultURL()))
		return
	}

	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO error page
		return
	}
	user, err := u.userService.Login(form.Username, form.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO error page
		return
	}

	sessionID, err := u.sessionService.CreateSession(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO error page
		return
	}

	c.SetCookie("session_id", sessionID, 60*60*24*30, "/", "", true, true)
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", u.config.GetDefaultURL()))
}

func (u *UserHandler) GetRegisterPage(c *gin.Context) {
	if u.isLoggedIn(c) {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", u.config.GetDefaultURL()))
		return
	}

	c.HTML(200, "register", gin.H{
		"Title":  "Register",
		"Header": u.config.GetHeader("", false),
	})
}

type RegisterForm struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email"`
	Password string `form:"password" binding:"required"`
}

func (u *UserHandler) Register(c *gin.Context) {
	if u.isLoggedIn(c) {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", u.config.GetDefaultURL()))
		return
	}

	var form RegisterForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO error page
		return
	}
	userType := u.config.GetNewUserType()
	err := u.userService.Register(form.Username, form.Password, form.Email, userType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // TODO error page
		return
	}

	u.Login(c)
}

func (u *UserHandler) Logout(c *gin.Context) {
	if !u.isLoggedIn(c) {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", u.config.GetDefaultURL()))
		return
	}

	cookie, _ := c.Cookie("session_id")
	u.sessionService.RemoveSession(cookie)

	c.SetCookie("session_id", "", 0, "/", "", true, true)
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s", u.config.GetDefaultURL()))
}

func (u *UserHandler) isLoggedIn(c *gin.Context) bool {
	cookie, err := c.Cookie("session_id")
	if err == nil {
		_, loggedIn := u.sessionService.IsUserLoggedIn(cookie)
		return loggedIn
	}
	return false
}
