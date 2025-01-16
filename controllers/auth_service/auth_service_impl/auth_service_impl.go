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

type NewAuthServiceImpl struct {
	Logger  logger.IAppLogger
	AuthDao dao.AuthDao
	DB      *sql.DB
}

func NewAuthService(input NewAuthServiceImpl) authservice.AuthService {
	return &AuthServiceImpl{
		logger:  input.Logger,
		authDao: input.AuthDao,
		db:      input.DB,
	}
}

func (a *AuthServiceImpl) OauthSetup(ctx *gin.Context, req *models.OauthUserInfo) (string, error) {

	return "", nil
}

func (a *AuthServiceImpl) SignUp(ctx *gin.Context, req *models.Users) (*models.Users, string, error) {
	a.logger.Info("Initiating SignUp process")

	exists, err := a.authDao.CheckUserExists(req.Email)
	if err != nil {

		return &models.Users{}, "error in checking if user exist", err
	}

	if exists {
		return &models.Users{}, "User already exists", nil
	}
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {

		return &models.Users{}, "Error hashing password", err
	}

	passwordMatch, msg := utils.VerifyPassword(hashed, req.Password)
	if !passwordMatch {
		a.logger.Info("Password hashing failed: ", msg)
		return &models.Users{}, "password is invalid", fmt.Errorf("password is invalid")
	}

	req.Password = hashed

	if req.CreatedAt.IsZero() {
		req.CreatedAt = time.Now().UTC()
	}

	userCreated, err := a.authDao.SignUp(req)
	if err != nil {
		a.logger.Error("Error during SignUp: ", err)
		return &models.Users{}, "Error during SignUp", err
	}

	return userCreated, "Signed Up successfully", nil
}

func (a *AuthServiceImpl) CheckUserExists(req *models.Users) (bool, error) {

	return a.authDao.CheckUserExists(req.Email)
}

func (a *AuthServiceImpl) ProcessLogin(ctx *gin.Context, req *models.LoginReq) (string, error) {

	exists, err := a.authDao.CheckUserExists(req.Email)
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
