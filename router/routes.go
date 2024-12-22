package router

import (
	"ecommerce-platform/middleware"
)

func (r *Router) SetupRoutes() {

	r.Engine.Use(middleware.EnableCors())

	adminGroup := r.Engine.Group("/admin")
	{
		adminGroup.Use(middleware.Auth([]string{"admin"}))
		adminGroup.GET("/health-check", r.HealthCheck)

	}

	clientGroup := r.Engine.Group("/client")
	{
		clientGroup.Use(middleware.Auth([]string{"client"}))
		clientGroup.GET("/health-check", r.HealthCheck)
	}

	userGroup := r.Engine.Group("/user")
	{
		userGroup.Use(middleware.Auth([]string{"user"}))
		userGroup.GET("/health-check", r.HealthCheck)
	}

	publicGroup := r.Engine.Group("/public")
	{

		publicGroup.GET("/health-check", r.HealthCheck)
		publicGroup.POST("/signup", r.SignUp)
	}
}
