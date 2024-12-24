package Dao

import "ecommerce-platform/models"

type AuthDao interface {
	SignUp(req *models.Users) (bool, error)
	CheckUserExistsSignup(req *models.Users) (bool, error)
	CheckUserExistsLogin(req *models.LoginReq) (bool, error)
	GetUser(req *models.LoginReq) (models.Users, error)
}

type AdminDao interface {
	CheckUserExists(req *models.Users) (bool, error)
}
