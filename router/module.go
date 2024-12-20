package router

import (
	validation "ecommerce-platform/Validation"
	authservice "ecommerce-platform/controllers/auth_service"
	log "ecommerce-platform/logger"
	adminroutes "ecommerce-platform/router/routes/admin_routes"
	clientroutes "ecommerce-platform/router/routes/client_routes"
	userroutes "ecommerce-platform/router/routes/user_routes"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine      *gin.Engine
	AuthService authservice.AuthService
	Logger      log.IAppLogger
	Val         validation.ValidationService
}

func NewRouter(
	logger log.IAppLogger,
	authService authservice.AuthService,
	valService validation.ValidationService,
) *Router {
	engine := gin.Default()
	router := &Router{
		Engine:      engine,
		AuthService: authService,
		Logger:      logger,
		Val:         valService,
	}
	adminroutes.AdminRoutes(router)
	clientroutes.ClientRoutes(router)
	userroutes.UserRoutes(router)
	return router
}
