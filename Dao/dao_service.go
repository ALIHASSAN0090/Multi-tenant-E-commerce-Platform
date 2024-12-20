package Dao

import (
	"database/sql"
	"ecommerce-platform/models"
)

type AuthDao interface {
	SignUp(req *models.Users) (bool, error)
	CheckUserExists(req *models.Users) (bool, error)
}

func NewAuthDao(db *sql.DB) AuthDao {
	return &AuthDaoImp{
		db: db,
	}
}
