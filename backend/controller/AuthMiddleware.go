package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-11/backend/secure"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := secure.GetAuthentication(c)
		if err != nil {
			c.JSON(401, gin.H{
				"status": 401,
				"error":  err.Error(),
			})
			c.SetCookie("token", "", -1, "/", "localhost", false, true)
			c.SetCookie("RLPP", "", -1, "/", "localhost", false, true)
			c.SetCookie("id", "", -1, "/", "localhost", false, true)
			c.Abort()
			return
		}
		res, err1 := secure.ExtractAuthToken(c)
		if err1 != nil {
			c.JSON(401, gin.H{
				"status": 401,
				"error":  err1.Error(),
			})
			c.SetCookie("token", "", -1, "/", "localhost", false, true)
			c.SetCookie("RLPP", "", -1, "/", "localhost", false, true)
			c.SetCookie("id", "", -1, "/", "localhost", false, true)
			c.Abort()
			return

		}
		id, _ := c.Cookie("id")
		newid, _ := strconv.Atoi(id)
		err2 := secure.Verifyid(newid, res.ID)
		if err2 != nil {
			c.JSON(401, gin.H{
				"status": 401,
				"error":  err2.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AuthMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("RLPP")
		if err != nil {
			c.JSON(401, gin.H{
				"status": 401,
				"error":  err.Error(),
			})
		}

		_, err1 := secure.VerifyCookie(cookie, "admin")
		if err1 != nil {
			c.JSON(401, gin.H{
				"error": err1.Error(),
			})
			c.Abort()
			return
		}

		c.Next()

	}
}
