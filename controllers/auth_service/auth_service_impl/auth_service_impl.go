package auth_service_impl

import (
	"database/sql"
	dao "ecommerce-platform/Dao"
	authservice "ecommerce-platform/controllers/auth_service"
	logger "ecommerce-platform/logger"
	"ecommerce-platform/middleware"
	"ecommerce-platform/models"
	"ecommerce-platform/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthServiceImpl struct {
	logger  logger.IAppLogger
	authDao dao.AuthDao
	db      *sql.DB
}

func NewAuthService(input NewAuthServiceImpl) authservice.AuthService {
	return &AuthServiceImpl{
		logger:  input.Logger,
		authDao: input.AuthDao,
		db:      input.DB,
	}
}

type NewAuthServiceImpl struct {
	Logger  logger.IAppLogger
	AuthDao dao.AuthDao
	DB      *sql.DB
}

func (a *AuthServiceImpl) SignUp(ctx *gin.Context, req *models.Users) (bool, error) {
	a.logger.Info("Initiating SignUp process")

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		a.logger.Error("Error hashing password: ", err)
		return false, err
	}

	req.Password = hashed

	exists, err := a.authDao.CheckUserExistsSignup(req)
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

	return a.authDao.CheckUserExistsSignup(req)
}

func (a *AuthServiceImpl) ProcessLogin(ctx *gin.Context, req *models.LoginReq) (string, error) {

	exists, err := a.authDao.CheckUserExistsLogin(req)
	if err != nil {
		a.logger.Error("Error checking if user exists: ", err)
		return "", err
	}

	if !exists {
		a.logger.Info("User does not exist")
		return "", fmt.Errorf("user does not exist")
	}

	user, err := a.authDao.GetUser(req)
	if err != nil {
		a.logger.Error("Error retrieving user: ", err)
		return "", err
	}

	passwordMatch, msg := utils.VerifyPassword(user.Password, req.Password)
	if !passwordMatch {
		a.logger.Info("Password verification failed: ", msg)
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := middleware.GenerateAccessToken(&user)
	utils.HandleError(err)
	return token, nil
}
