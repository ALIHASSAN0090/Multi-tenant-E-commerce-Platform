package Dao

import (
	"database/sql"
)

type AuthDao interface {
	// SignUp(req *models.Users) (bool, error)
	// CheckUserExists(req *models.Users) (bool, error)
}

type AuthDaoImp struct {
	db *sql.DB
}

func NewAuthDao(db *sql.DB) AuthDao {
	return &AuthDaoImp{
		db: db,
	}
}
