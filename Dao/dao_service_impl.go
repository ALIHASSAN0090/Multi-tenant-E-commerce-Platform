package Dao

import (
	"database/sql"
	"ecommerce-platform/models"
)

type AuthDaoImp struct {
	db *sql.DB
}

func (a *AuthDaoImp) CheckUserExists(req *models.Users) (bool, error) {
	var exists bool

	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND role = 'user' )`
	err := a.db.QueryRow(checkQuery, req.Email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (a *AuthDaoImp) SignUp(req *models.Users) (bool, error) {

	query := `INSERT INTO users(user_name, email, password, phone_number, role, address, created_at, updated_at) 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
	stmt, err := a.db.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.UserName, req.Email, req.Password, req.PhoneNumber, req.Role, req.Address, req.CreatedAt, req.UpdatedAt)
	if err != nil {
		return false, err
	}

	return false, nil
}
