package seller_controller

import "ecommerce-platform/models"

type SellerController interface {
	GetStoreItems(seller_id int64) ([]models.Item, error)
}
