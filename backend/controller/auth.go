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
	authService  service.AuthService
	userService  service.UserService
	bookService  service.BookService
	adminService service.AdminService
}

func NewAuthHandler(authService service.AuthService, userService service.UserService, bookService service.BookService, adminService service.AdminService) *AuthHandler {
	return &AuthHandler{authService, userService, bookService, adminService}
}
func (a *AuthHandler) Login(c *gin.Context) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
	var loginReq model.PayloadUser
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	Role, err := a.userService.GetRoleByUserName(loginReq.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  err.Error(),
		})
		return
	}
	id, err := a.userService.GetIdByUserName(loginReq.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"status": 400,
			"error":  err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
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
	var data model.Token
	data.Token = token
	data.Role = hashcookie
	data.ID = id
	c.Writer.Header().Set("Authorization", token)
	c.Writer.Header().Set("RLPP", hashcookie)
	c.Writer.Header().Set("id", strconv.Itoa(id))

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Login Success",
		"data":    data,
	})
}

func (a *AuthHandler) Register(c *gin.Context) {
	config.Mutex.Lock()
	defer config.Mutex.Unlock()
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
			"status": 400,
			"error":  err.Error(),
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
	c.JSON(200, gin.H{
		"status":  200,
		"message": "Logout Success",
	})
}

type Token struct {
	Id int `json:"id"`
}

// func (a *AuthHandler) RefreshToken(c *gin.Context) {
// 	var token Token
// 	if err := c.ShouldBindJSON(&token); err != nil {
// 		c.JSON(400, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	token1, err := secure.RefreshToken(token.Id)

// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"status":  200,
// 		"message": "Refresh Token Success",
// 		"token":   token1,
// 	})
// }
