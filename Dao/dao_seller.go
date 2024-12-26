package Dao

import "database/sql"

type SellerDaoImpl struct {
	db *sql.DB
}

func NewSellerDao(db *sql.DB) SellerDao {
	return &SellerDaoImpl{
		db: db,
	}

}
