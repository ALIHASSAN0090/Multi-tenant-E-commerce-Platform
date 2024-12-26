package Dao

import (
	"database/sql"
	"ecommerce-platform/models"
)

type SellerDaoImpl struct {
	db *sql.DB
}

func NewSellerDao(db *sql.DB) *SellerDaoImpl {
	return &SellerDaoImpl{
		db: db,
	}
}

func (s *SellerDaoImpl) GetStoreItemsDB(sellerID int64) ([]models.Item, error) {
	storeID, err := s.GetStoreIDByUserID(sellerID)
	if err != nil {
		return nil, err
	}

	query := `
	SELECT id, name, price, stock_quantity, description, discount 
	FROM items 
	WHERE store_id = $1
	`
	rows, err := s.db.Query(query, storeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.StockQuantity, &item.Description, &item.Discount); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *SellerDaoImpl) GetStoreIDByUserID(userID int64) (int64, error) {
	query := `
	SELECT st.id AS store_id
	FROM seller AS s
	JOIN stores st ON st.seller_id = s.id
	WHERE s.user_id = $1
	`

	var storeID int64
	err := s.db.QueryRow(query, userID).Scan(&storeID)
	if err != nil {
		return 0, err
	}

	return storeID, nil
}
