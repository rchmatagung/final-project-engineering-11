package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	app.Run(":8080")
}
