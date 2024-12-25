package router

import (
	"ecommerce-platform/Validation"
	"ecommerce-platform/controllers/admin_controller"
	"ecommerce-platform/controllers/auth_service"
	dao_service_impl "ecommerce-platform/logger"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine      *gin.Engine
	AuthService auth_service.AuthService
	Logger      dao_service_impl.IAppLogger
	Val         Validation.ValidationService
	Admin       admin_controller.AdminControllers
}

func NewRouter(
	logger dao_service_impl.IAppLogger,
	authService auth_service.AuthService,
	valService Validation.ValidationService,
	AdminController admin_controller.AdminControllers,
) *Router {

	engine := gin.Default()

	router := &Router{
		Engine:      engine,
		AuthService: authService,
		Logger:      logger,
		Val:         valService,
		Admin:       AdminController,
	}
	router.SetupRoutes()
	return router
}
