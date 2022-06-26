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
			c.Abort()
			return
		}
		res, err1 := secure.ExtractAuthToken(c)
		if err1 != nil {
			c.JSON(401, gin.H{
				"status": 401,
				"error":  err1.Error(),
			})
			c.Abort()
			return

		}
		id := c.GetHeader("id")
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
		cookie := c.GetHeader("RLPP")
		if cookie == "" {
			c.JSON(401, gin.H{
				"status": 401,
				"error":  "Unauthorized",
			})
			c.Abort()
			return
		}

		_, err1 := secure.VerifyCookie(cookie, "admin")
		if err1 != nil {
			c.JSON(401, gin.H{
				"status": 401,
				"error":  err1.Error(),
			})
			c.Abort()
			return
		}

		c.Next()

	}
}
