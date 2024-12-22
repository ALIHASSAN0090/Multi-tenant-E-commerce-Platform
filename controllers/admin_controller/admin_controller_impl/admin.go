package admin_controller_impl

import (
	dao "ecommerce-platform/Dao"
	logger "ecommerce-platform/logger"
	"ecommerce-platform/models"

	admin_controller "ecommerce-platform/controllers/admin_controller"

	"github.com/gin-gonic/gin"
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

func (a *AdminControllerImpl) SignUp(ctx *gin.Context, req *models.Users) (bool, error) {
	return false, nil
}
