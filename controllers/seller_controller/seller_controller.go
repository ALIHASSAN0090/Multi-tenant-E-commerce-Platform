package seller_controller

import (
	"ecommerce-platform/models"

	"github.com/gin-gonic/gin"
)

type SellerController interface {
	GetStoreItems(seller_id int64) ([]models.Item, error)
	GetStoreItem(c *gin.Context, id int64) (models.Item, error)
	UpdateStoreItem(c *gin.Context, id int64, item models.Item) (models.Item, error)
}
