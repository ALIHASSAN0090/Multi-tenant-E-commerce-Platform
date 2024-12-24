package Dao

import "ecommerce-platform/models"

type AuthDao interface {
	SignUp(req *models.Users) (*models.Users, error)

	CheckUserExists(req string) (bool, error)
	GetUser(req *models.LoginReq) (models.Users, error)
}

type AdminDao interface {
}
