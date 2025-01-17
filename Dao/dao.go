package Dao

import (
	"database/sql"
	"ecommerce-platform/models"

	"github.com/gin-gonic/gin"
)

type AuthDao interface {
	SignUp(req *models.Users) (*models.Users, error)
	CheckUserExists(req string) (bool, error)
	GetUser(req *models.LoginReq) (models.Users, error)
}

type AdminDao interface {
}

type UserDao interface {
	CreateSeller(seller models.SellerStore) (models.Seller, error)
	CreateStore(store models.SellerStore, seller_id int64) (models.Store, error)
	checkExistingSeller(userID int64) (bool, error)
	ChangeRoleToSeller(id int64) (bool, error)
	GetStores() ([]models.Store, error)
	GetStoreItems(store_id int64) (models.StoreItems, error)
	GetTotalPriceUnitPrice(tx *sql.Tx, items []models.OrderItem) (float64, []int64, []float64, error)
	CreateOrder(tx *sql.Tx, orderData models.CreateOrder) (models.Order, error)
	CreateItems(tx *sql.Tx, items []models.OrderItem) ([]models.OrderItem, error)
}

type SellerDao interface {
	GetStoreItemsDB(seller_id int64) ([]models.Item, error)
	GetStoreItemDB(id int64) (models.Item, error)
	UpdateStoreItem(id int64, item models.Item) (models.Item, error)
	GetStoreIDByUserID(userID int64) (int64, error)
	CreateItem(store_id int64, item models.Item) (models.Item, error)
	GetStore(sellerID int64) (models.Store, error)
	IsActive(c *gin.Context, seller_id int64) (bool, error)
	GetAllOrders(store_id int64, filter string) ([]models.Order, error)
	GetOrderByOrderId(order_id int64) (models.Order, error)
	GetCustomerNameByOrderId(order_id int64) (string, error)
	GetOrderItemsByOrderId(order_id int64) ([]models.ItemResponce, error)
}
