package userroutes

import (
	"ecommerce-platform/middleware"
	"ecommerce-platform/router"
)

func UserRoutes(r *router.Router) {
	r.Engine.Use(middleware.EnableCors())
	// Add your admin-specific routes here
	// r.Engine.GET("/health-check", r.HealthCheck)
	// r.Engine.POST("/signup", r.SignUp)
}
