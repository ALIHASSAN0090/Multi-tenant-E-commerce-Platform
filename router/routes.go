package router

import (
	"ecommerce-platform/middleware"
)

func (r *Router) SetupRoutes() {
	rateLimiter := middleware.NewRateLimiter(1, 5)
	r.Engine.Use(middleware.EnableCors(), rateLimiter.Limit())

	adminGroup := r.Engine.Group("/admin")
	{
		adminGroup.Use(middleware.Auth([]string{"admin"}))
		adminGroup.GET("/health-check", r.HealthCheck)

	}

	sellerGroup := r.Engine.Group("/seller")
	{
		sellerGroup.Use(middleware.Auth([]string{"seller", "admin"}))
		sellerGroup.GET("/health-check", r.HealthCheck)
		sellerGroup.GET("/items", r.GetStoreItems)

	}

	userGroup := r.Engine.Group("/user")
	{
		userGroup.Use(middleware.Auth([]string{"user", "admin", "seller"}))
		userGroup.GET("/health-check", r.HealthCheck)
		userGroup.POST("/create/seller/store", r.CreateSeller)
	}

	publicGroup := r.Engine.Group("/public")
	{

		publicGroup.GET("/health-check", r.HealthCheck)
		publicGroup.POST("/signup", r.SignUp)
		publicGroup.POST("/login", r.Login)
	}
}
