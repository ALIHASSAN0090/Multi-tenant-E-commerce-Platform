package Dao

import (
	"database/sql"
	"ecommerce-platform/models"
	"fmt"
	"time"
)

func NewUserDao(db *sql.DB) UserDao {
	return &UserDaoImpl{
		db: db,
	}
}

type UserDaoImpl struct {
	db *sql.DB
}

func (dao *UserDaoImpl) CreateSeller(seller models.SellerStore) (models.Seller, error) {

	exists, err := dao.checkExistingSeller(seller.UserID)
	if err != nil {
		return models.Seller{}, err
	}
	if exists {
		return models.Seller{}, fmt.Errorf("seller with user_id %d already exists", seller.UserID)
	}

	query := `INSERT INTO seller(user_id, business_name, contact_number, created_at, active, seller_img, cnic_number, cnic_image) 
	values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	var sellerID int64
	err = dao.db.QueryRow(query, seller.UserID, seller.BusinessName, seller.ContactNumber, time.Now(), true, seller.SellerImg, seller.CNICNumber, seller.CNICImage).Scan(&sellerID)
	if err != nil {
		return models.Seller{}, err
	}

	return models.Seller{
		ID:            sellerID,
		UserID:        seller.UserID,
		BusinessName:  seller.BusinessName,
		ContactNumber: seller.ContactNumber,
		CreatedAt:     time.Now(),
		Active:        seller.Active,
		SellerImg:     seller.SellerImg,
		CNICNumber:    seller.CNICNumber,
		CNICImage:     seller.CNICImage,
	}, nil
}

func (dao *UserDaoImpl) CreateStore(store models.SellerStore, seller_id int64) (models.Store, error) {
	query := `INSERT INTO stores(seller_id, store_img, store_name, store_address, store_description, created_at) 
	values ($1, $2, $3, $4, $5, $6) returning id`

	var storeID int64
	err := dao.db.QueryRow(query, seller_id, store.StoreImg, store.StoreName, store.StoreAddress, store.StoreDescription, time.Now()).Scan(&storeID)
	if err != nil {
		return models.Store{}, err
	}

	return models.Store{
		ID:               storeID,
		SellerID:         seller_id,
		StoreImg:         store.StoreImg,
		StoreName:        store.StoreName,
		StoreAddress:     store.StoreAddress,
		StoreDescription: store.StoreDescription,
		CreatedAt:        time.Now(),
	}, nil
}

func (dao *UserDaoImpl) checkExistingSeller(userID int64) (bool, error) {
	existingSellerQuery := `SELECT id FROM seller WHERE user_id = $1`
	var existingSellerID int64
	err := dao.db.QueryRow(existingSellerQuery, userID).Scan(&existingSellerID)
	if err == nil {
		return true, fmt.Errorf("user_id already exists in the seller table")
	} else if err != sql.ErrNoRows {
		return false, err
	}
	return false, nil
}
