package Dao

import (
	"database/sql"
)

func NewAdminDao(db *sql.DB) AdminDao {
	return &AdminDaoImpl{
		db: db,
	}
}

type AdminDaoImpl struct {
	db *sql.DB
}
