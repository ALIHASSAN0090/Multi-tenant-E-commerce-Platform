package auth_service

import (
	"ecommerce-platform/models"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	SignUp(ctx *gin.Context, req *models.Users) (bool, error)
}
