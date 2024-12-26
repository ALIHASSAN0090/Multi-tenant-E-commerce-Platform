package router

import (
	"ecommerce-platform/Validation"
	"ecommerce-platform/controllers/admin_controller"
	"ecommerce-platform/controllers/auth_service"
	"ecommerce-platform/controllers/seller_controller"
	"ecommerce-platform/controllers/user_controller"
	logger "ecommerce-platform/logger"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine           *gin.Engine
	AuthService      auth_service.AuthService
	Logger           logger.IAppLogger
	Val              Validation.ValidationService
	Admin            admin_controller.AdminControllers
	UserController   user_controller.UserControllerConfig
	SellerController seller_controller.SellerController
}

func NewRouter(
	logger logger.IAppLogger,
	authService auth_service.AuthService,
	valService Validation.ValidationService,
	AdminController admin_controller.AdminControllers,
	userControllers user_controller.UserControllerConfig,
	sellerController seller_controller.SellerController,
) *Router {

	engine := gin.Default()

	router := &Router{
		Engine:           engine,
		AuthService:      authService,
		Logger:           logger,
		Val:              valService,
		Admin:            AdminController,
		UserController:   userControllers,
		SellerController: sellerController,
	}
	router.SetupRoutes()
	return router
}
