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

func (a *AuthDaoImp) CheckUserExists(req string) (bool, error) {

	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	err := a.db.QueryRow(checkQuery, req).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (a *AuthDaoImp) SignUp(req *models.Users) (*models.Users, error) {

	query1 := `INSERT INTO users(name, email, hash_password, role_id, created_at) 
	VALUES($1, $2, $3, $4, NOW()) RETURNING id`
	stmt, err := a.db.Prepare(query1)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var roleID int
	query2 := `SELECT id FROM roles WHERE name = 'user'`
	err = a.db.QueryRow(query2).Scan(&roleID)
	if err != nil {
		return nil, err
	}

	if err := stmt.QueryRow(req.UserName, req.Email, req.Password, roleID).Scan(&req.ID); err != nil {
		return nil, err
	}

	return req, nil
}

func (a *AuthDaoImp) GetUser(req *models.LoginReq) (models.Users, error) {
	var user models.Users

	query := `
	SELECT u.id, u.name, u.email, u.hash_password, r.name as role, u.created_at
    FROM users u
    Join roles r
    on r.id = u.role_id
    WHERE email = $1
	`
	err := a.db.QueryRow(query, req.Email).Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}
