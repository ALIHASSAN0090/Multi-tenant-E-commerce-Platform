package Dao

import (
	"database/sql"
	"ecommerce-platform/Dao/dao_service_impl"
	"ecommerce-platform/models"
)

type AuthDao interface {
	SignUp(req *models.Users) (bool, error)
	CheckUserExists(req *models.Users) (bool, error)
}

func NewAuthDao(db *sql.DB) AuthDao {
	return &dao_service_impl.AuthDaoImp{
		db: db,
	}
}
