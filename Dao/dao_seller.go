package Dao

import (
	"database/sql"
	"ecommerce-platform/models"
	"fmt"
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
	if err != nil || storeID == 0 {
		return nil, fmt.Errorf("invalid store ID for seller ID: %d", sellerID)
	}

	query := `
	SELECT id, name, store_id,price, stock_quantity,item_img, description, discount 
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
		if err := rows.Scan(&item.ID, &item.Name, &item.StoreID, &item.Price, &item.StockQuantity, &item.ItemImg, &item.Description, &item.Discount); err != nil {
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

func (s *SellerDaoImpl) GetStoreItemDB(id int64) (models.Item, error) {
	query := `
	SELECT id, name, store_id, price, stock_quantity, item_img, description, discount 
	FROM items 
	WHERE id = $1
	`

	var item models.Item
	err := s.db.QueryRow(query, id).Scan(&item.ID, &item.Name, &item.StoreID, &item.Price, &item.StockQuantity, &item.ItemImg, &item.Description, &item.Discount)
	if err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func (s *SellerDaoImpl) UpdateStoreItem(id int64, item models.Item) (models.Item, error) {

	query := `
		UPDATE items
		SET name = $1, price = $2, stock_quantity = $3, item_img = $4, description = $5, discount = $6, updated_at = NOW()
		WHERE id = $7
		RETURNING id, name, store_id, price, stock_quantity, item_img, description, discount, updated_at
	`

	var updatedItem models.Item
	err := s.db.QueryRow(query, item.Name, item.Price, item.StockQuantity, item.ItemImg, item.Description, item.Discount, id).Scan(
		&updatedItem.ID, &updatedItem.Name, &updatedItem.StoreID, &updatedItem.Price, &updatedItem.StockQuantity, &updatedItem.ItemImg, &updatedItem.Description, &updatedItem.Discount, &updatedItem.UpdatedAt,
	)
	if err != nil {
		return models.Item{}, err
	}

	return updatedItem, nil
}
