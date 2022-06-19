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
			c.SetCookie("token", "", -1, "/", "localhost", false, true)
			c.SetCookie("RLPP", "", -1, "/", "localhost", false, true)
			c.SetCookie("id", "", -1, "/", "localhost", false, true)
			c.JSON(401, gin.H{
				"status": 401,
				"error":  err.Error(),
			})
			c.Abort()
			return
		}
		res, err1 := secure.ExtractAuthToken(c)
		if err1 != nil {
			c.SetCookie("token", "", -1, "/", "localhost", false, true)
			c.SetCookie("RLPP", "", -1, "/", "localhost", false, true)
			c.SetCookie("id", "", -1, "/", "localhost", false, true)
			c.JSON(401, gin.H{
				"status": 401,
				"error":  err1.Error(),
			})
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
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
				"status": 401,
				"error":  err1.Error(),
			})
			c.Abort()
			return
		}

		c.Next()

	}
}
