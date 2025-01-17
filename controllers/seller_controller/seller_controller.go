package seller_controller

import (
	"ecommerce-platform/models"

	"github.com/gin-gonic/gin"
)

type SellerController interface {
	GetStoreItems(seller_id int64) ([]models.Item, error)
	GetStoreItem(c *gin.Context, id int64) (models.Item, error)
	UpdateStoreItem(c *gin.Context, id int64, item models.Item) (models.Item, error)
	CreateItem(c *gin.Context, seller_id int64, item models.Item) (models.Item, error)
	GetStore(c *gin.Context, seller_id int64) (models.Store, error)
	GetAllOrders(c *gin.Context, user_id int64, filter string) ([]models.Order, error)
	GetOrderByOrderId(c *gin.Context, order_id int64) (models.OrderResponce, error)
}
