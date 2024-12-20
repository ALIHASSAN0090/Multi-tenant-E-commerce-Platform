package router

import (
	"ecommerce-platform/middleware"
)

func (r *Router) SetupRoutes() {

	r.Engine.Use(middleware.EnableCors())

	adminGroup := r.Engine.Group("/admin")
	{

		adminGroup.GET("/health-check", r.SignUp)

	}

	clientGroup := r.Engine.Group("/client")
	{
		clientGroup.GET("/health-check", r.SignUp)
	}

	userGroup := r.Engine.Group("/user")
	{
		userGroup.GET("/health-check", r.SignUp)
	}

	publicGroup := r.Engine.Group("/public")
	{

		// publicGroup.POST("/login", r.Login)
		publicGroup.POST("/signup", r.SignUp)
	}
}
