package user_controller

import (
	"ecommerce-platform/models"

	"github.com/gin-gonic/gin"
)

type UserControllerConfig interface {
	CreateSellerStore(c *gin.Context, seller models.SellerStore) (models.Seller, models.Store, error)
	GetStores(c *gin.Context) ([]models.Store, error)
}
