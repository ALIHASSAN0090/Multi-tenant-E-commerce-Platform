package admin_controller

import (
	"ecommerce-platform/models"

	"github.com/gin-gonic/gin"
)

type AdminControllers interface {
	SignUp(ctx *gin.Context, req *models.Users) (bool, error)
}
