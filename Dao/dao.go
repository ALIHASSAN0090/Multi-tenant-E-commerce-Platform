package Dao

import "ecommerce-platform/models"

type AuthDao interface {
	SignUp(req *models.Users) (*models.Users, error)
	CheckUserExists(req string) (bool, error)
	GetUser(req *models.LoginReq) (models.Users, error)
}

type AdminDao interface {
}

type UserDao interface {
	CreateSeller(models.SellerStore) (models.Seller, error)
	CreateStore(store models.SellerStore, seller_id int64) (models.Store, error)
	checkExistingSeller(userID int64) (bool, error)
	ChangeRoleToSeller(id int64) (bool, error)
}

type SellerDao interface {
	GetStoreItemsDB(seller_id int64) ([]models.Item, error)
}
