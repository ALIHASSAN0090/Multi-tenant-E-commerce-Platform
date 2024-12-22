package Dao

import "ecommerce-platform/models"

type AuthDao interface {
	SignUp(req *models.Users) (bool, error)
	CheckUserExists(req *models.Users) (bool, error)
}

type AdminDao interface {
	CheckUserExists(req *models.Users) (bool, error)
}
