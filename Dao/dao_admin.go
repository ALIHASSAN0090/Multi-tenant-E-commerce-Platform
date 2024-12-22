package Dao

import (
	"database/sql"
	"ecommerce-platform/models"
)

func NewAdminDao(db *sql.DB) AdminDao {
	return &AdminDaoImpl{
		db: db,
	}
}

type AdminDaoImpl struct {
	db *sql.DB
}

func (a *AdminDaoImpl) CheckUserExists(req *models.Users) (bool, error) {
	var exists bool

	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND role = 'user' )`
	err := a.db.QueryRow(checkQuery, req.Email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
