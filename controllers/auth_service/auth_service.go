package auth_service

import (
	"ecommerce-platform/models"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	SignUp(ctx *gin.Context, req *models.Users) (*models.Users, string, error)
	CheckUserExists(req *models.Users) (bool, error)
	ProcessLogin(ctx *gin.Context, req *models.LoginReq) (string, error)
	OauthSetup(ctx *gin.Context, req *models.OauthUserInfo) (string, error)
}
