package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-11/backend/config"
	"github.com/rg-km/final-project-engineering-11/backend/model"
	"github.com/rg-km/final-project-engineering-11/backend/secure"
	"github.com/rg-km/final-project-engineering-11/backend/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}
func (a *AuthHandler) Login(c *gin.Context) {
	var loginReq model.PayloadUser
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	Role, err := a.authService.GetRoleByUserName(loginReq.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  err.Error(),
		})
		return
	}
	id, err := a.authService.GetIdByUserName(loginReq.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  err.Error(),
		})
		return
	}

	token, err := a.authService.Login(loginReq)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashcookie, _ := secure.HashPassword(Role)
	c.SetCookie("token", token, int(config.Configuration.JWT_EXPIRATION_DURATION.Seconds()), "/", "localhost", false, true)
	c.SetCookie("id", strconv.Itoa(id), int(config.Configuration.JWT_EXPIRATION_DURATION.Seconds()), "/", "localhost", false, true)
	c.SetCookie("RLPP", hashcookie, int(config.Configuration.JWT_EXPIRATION_DURATION.Seconds()), "/", "localhost", false, true)

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Login Success",
		"token":   token,
	})
}

func (a *AuthHandler) Register(c *gin.Context) {
	var register model.UserRegis
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := a.authService.Register(register)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Register Success",
	})
}

func (a *AuthHandler) Logout(c *gin.Context) {
	c.Header("Authorization", "")
	c.SetCookie("RLPP", "", -1, "/", "localhost", false, true)
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.SetCookie("id", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Logout Success",
	})
}
