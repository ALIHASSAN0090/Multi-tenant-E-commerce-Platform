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
		sellerGroup.Use(middleware.StatusCheck(r.Engine.HandleContext))

		sellerGroup.GET("/health-check", r.HealthCheck)
		sellerGroup.GET("/items", r.GetStoreItems)
		sellerGroup.GET("/item/:id", r.GetStoreItem)
		sellerGroup.POST("/item", r.CreateItem)
		sellerGroup.PATCH("/item/:id", r.UpdateItem)
		sellerGroup.GET("/store", r.GetStore)

	}

	userGroup := r.Engine.Group("/user")
	{
		userGroup.Use(middleware.Auth([]string{"user", "admin", "seller"}))

		userGroup.GET("/health-check", r.HealthCheck)
		userGroup.POST("/create/seller/store", r.CreateSeller)
		userGroup.GET("/stores", r.GetStores)
		userGroup.GET("/store/:id", r.GetStoreAndItems)
		userGroup.GET("/store/item/:id", r.GetStoreItem)
		userGroup.POST("/create/order", r.CreateOrder)
	}

	publicGroup := r.Engine.Group("/public")
	{

		publicGroup.GET("/health-check", r.HealthCheck)
		publicGroup.POST("/signup", r.SignUp)
		publicGroup.POST("/login", r.Login)
	}
}
