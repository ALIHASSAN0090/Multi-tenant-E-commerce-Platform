package auth_service_impl

import (
	"ecommerce-platform/models"
	"ecommerce-platform/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *AuthServiceImpl) SignUp(ctx *gin.Context, req *models.Users) (bool, error) {
	a.logger.Info("Initiating SignUp process")

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		a.logger.Error("Error hashing password: ", err)
		return false, err
	}

	req.Password = hashed

	exists, err := a.authDao.CheckUserExists(req)
	if err != nil {
		a.logger.Error("Error checking if user exists: ", err)
		return false, err
	}

	if exists {
		a.logger.Info("User already exists")
		return true, nil
	}

	if req.CreatedAt.IsZero() {
		req.CreatedAt = time.Now().UTC()
	}
	req.UpdatedAt = time.Now().UTC()

	userExists, err := a.authDao.SignUp(req)
	if err != nil {
		a.logger.Error("Error during SignUp: ", err)
		return false, err
	}

	return userExists, nil
}

func (a *AuthServiceImpl) CheckUserExists(req *models.Users) (bool, error) {
	return a.authDao.CheckUserExists(req)
}
