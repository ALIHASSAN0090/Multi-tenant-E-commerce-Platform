package seller_controller_impl

import (
	"database/sql"
	"ecommerce-platform/Dao"
	"ecommerce-platform/controllers/seller_controller"
	"ecommerce-platform/models"
)

type SellerControllerImpl struct {
	SellerDao Dao.SellerDao
	db        *sql.DB
}

type SellerController struct {
	SellerDao Dao.SellerDao
	DB        *sql.DB
}

func NewSellerImpl(input SellerController) seller_controller.SellerController {
	return &SellerControllerImpl{
		SellerDao: input.SellerDao,
		db:        input.DB,
	}
}

func (sc *SellerControllerImpl) GetStore(seller_id int64) (models.GetStoreData, error) {

}
