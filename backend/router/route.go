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
	api.Use(controller.CORSMiddleware())
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authentication.Login)
			auth.POST("/register", authentication.Register)
			auth.POST("/logout", authentication.Logout)

		}

		auth2 := api.Group("/user").Use(controller.AuthMiddleware())
		{
			auth2.GET("/mentor/detail/:id", authentication.DetailMentor)
			auth2.PUT("/update/:id", authentication.UpdateUserById)
			auth2.GET("/profile", authentication.UserProfile)
			auth2.GET("/booking/mentor/:id", authentication.GetRequestMentoring)
			auth2.GET("/booking/status", authentication.GetAllBookStatusMember)
			auth2.GET("/mentor/mentorlist", authentication.GetMentor)

		}
		auth3 := api.Group("/admin").Use(controller.AuthMiddleware()).Use(controller.AuthMiddlewareAdmin())
		{
			auth3.POST("/registermentor", authentication.RegisMentor)
			auth3.POST("/addartikel", authentication.CreateArtikel)

		}
		router.route.LoadHTMLFiles("template/accepted.tmpl", "template/errro.tmpl")
		router.route.GET("/api/mentor/acc/:bookid", authentication.UpdateSatusBooking)

	}

	return router
}

func (a *Router) Run(port string) {
	a.route.Run(port)
}
