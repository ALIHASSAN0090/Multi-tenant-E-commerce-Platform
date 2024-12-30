package Dao

import (
	"database/sql"
	"ecommerce-platform/models"
	"fmt"

	"github.com/gin-gonic/gin"
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

func (s *SellerDaoImpl) CreateItem(store_id int64, item models.Item) (models.Item, error) {

	checkQuery := `
	SELECT COUNT(*)
	FROM items
	WHERE name = $1 AND store_id = $2
	`

	var count int
	err := s.db.QueryRow(checkQuery, item.Name, store_id).Scan(&count)
	if err != nil {
		return models.Item{}, err
	}

	if count > 0 {
		return models.Item{}, fmt.Errorf("item with name '%s' already exists in store", item.Name)
	}

	query := `
	INSERT INTO items (name, store_id, price, stock_quantity, item_img, description, discount, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	RETURNING id, name, store_id, price, stock_quantity, item_img, description, discount, created_at
	`

	var createdItem models.Item
	err = s.db.QueryRow(query, item.Name, store_id, item.Price, item.StockQuantity, item.ItemImg, item.Description, item.Discount).Scan(
		&createdItem.ID, &createdItem.Name, &createdItem.StoreID, &createdItem.Price, &createdItem.StockQuantity, &createdItem.ItemImg, &createdItem.Description, &createdItem.Discount, &createdItem.CreatedAt,
	)
	if err != nil {
		return models.Item{}, err
	}

	return createdItem, nil
}

func (s *SellerDaoImpl) GetStore(sellerID int64) (models.Store, error) {
	query := `
	SELECT id, store_img, seller_id, store_name, store_description, store_address, created_at, updated_at 
	FROM stores 
	WHERE seller_id = $1
	`

	var store models.Store
	err := s.db.QueryRow(query, sellerID).Scan(
		&store.ID, &store.StoreImg, &store.SellerID, &store.StoreName, &store.StoreDescription, &store.StoreAddress, &store.CreatedAt, &store.UpdatedAt,
	)
	if err != nil {
		return models.Store{}, err
	}

	return store, nil
}

func (s *SellerDaoImpl) IsActive(c *gin.Context, seller_id int64) (bool, error) {
	query := `
	SELECT active
	FROM seller
	WHERE id = $1
	`

	var active bool
	err := s.db.QueryRow(query, seller_id).Scan(&active)
	if err != nil {
		return false, err
	}

	return active, nil
}
