package auth_service_impl

import (
	dao "ecommerce-platform/Dao"
	authservice "ecommerce-platform/controllers/auth_service"
	logger "ecommerce-platform/logger"
)

type AuthServiceImpl struct {
	logger  logger.IAppLogger
	authDao dao.AuthDao
}

func NewAuthService(input NewAuthServiceImpl) authservice.AuthService {
	return &AuthServiceImpl{
		logger:  input.Logger,
		authDao: input.AuthDao,
	}
}

type NewAuthServiceImpl struct {
	Logger  logger.IAppLogger
	AuthDao dao.AuthDao
}
