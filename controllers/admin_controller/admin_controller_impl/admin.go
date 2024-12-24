package admin_controller_impl

import (
	dao "ecommerce-platform/Dao"
	logger "ecommerce-platform/logger"

	admin_controller "ecommerce-platform/controllers/admin_controller"
)

type AdminControllerImpl struct {
	logger   logger.IAppLogger
	authDao  dao.AuthDao
	adminDao dao.AdminDao
}

func NewAdminController(input NewAdminControllerImpl) admin_controller.AdminControllers {
	return &AdminControllerImpl{
		logger:   input.Logger,
		authDao:  input.AuthDao,
		adminDao: input.AdminDao,
	}
}

type NewAdminControllerImpl struct {
	Logger   logger.IAppLogger
	AuthDao  dao.AuthDao
	AdminDao dao.AdminDao
}
