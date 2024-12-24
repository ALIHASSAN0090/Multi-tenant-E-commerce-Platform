package Dao

import (
	"database/sql"
	"ecommerce-platform/models"
)

func NewAuthDao(db *sql.DB) AuthDao {
	return &AuthDaoImp{
		db: db,
	}
}

type AuthDaoImp struct {
	db *sql.DB
}

func (a *AuthDaoImp) CheckUserExistsSignup(req *models.Users) (bool, error) {
	var exists bool

	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND role = 'user' )`
	err := a.db.QueryRow(checkQuery, req.Email).Scan(&exists)
	if err != nil {

		return false, err
	}

	return exists, nil
}

func (a *AuthDaoImp) CheckUserExistsLogin(req *models.LoginReq) (bool, error) {

	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
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

func (a *AuthDaoImp) GetUser(req *models.LoginReq) (models.Users, error) {
	var user models.Users

	query := `
	SELECT u.name, u.email, u.hash_password, r.name as role, u.created_at, u.updated_at 
    FROM users u
    Join roles r
    on r.id = u.role_id
    WHERE email = $1
	`
	err := a.db.QueryRow(query, req.Email).Scan(&user.UserName, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}
