package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-11/backend/controller"
)

type Router struct {
	route *gin.Engine
}

func Newrouter(authentication *controller.AuthHandler) *Router {
	router := &Router{
		route: gin.Default(),
	}

	api := router.route.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authentication.Login)
			auth.POST("/register", authentication.Register)
			auth.POST("/logout", authentication.Logout)
		}

		auth2 := api.Group("/user").Use(controller.AuthMiddleware())
		{
			auth2.GET("/all", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Hello World"}) })

		}

	}

	return router
}

func (a *Router) Run(port string) {
	a.route.Run(port)
}
