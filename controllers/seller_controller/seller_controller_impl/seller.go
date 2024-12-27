package seller_controller_impl

import (
	"database/sql"
	"ecommerce-platform/Dao"
	"ecommerce-platform/controllers/seller_controller"
	"ecommerce-platform/models"

	"github.com/gin-gonic/gin"
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

func (sc *SellerControllerImpl) GetStoreItems(seller_id int64) ([]models.Item, error) {

	return sc.SellerDao.GetStoreItemsDB(seller_id)

}

func (sc *SellerControllerImpl) GetStoreItem(c *gin.Context, id int64) (models.Item, error) {

	return sc.SellerDao.GetStoreItemDB(id)

}

func (sc *SellerControllerImpl) UpdateStoreItem(c *gin.Context, id int64, item models.Item) (models.Item, error) {

	return sc.SellerDao.UpdateStoreItem(id, item)

}

func (sc *SellerControllerImpl) CreateItem(c *gin.Context, seller_id int64, item models.Item) (models.Item, error) {
	store_id, err := sc.SellerDao.GetStoreIDByUserID(seller_id)
	if err != nil {
		return models.Item{}, err
	}

	createdItem, err := sc.SellerDao.CreateItem(store_id, item)
	if err != nil {
		return models.Item{}, err
	}

	return createdItem, nil
}

func (sc *SellerControllerImpl) GetStore(c *gin.Context, seller_id int64) (models.Store, error) {

	store_id, err := sc.SellerDao.GetStoreIDByUserID(seller_id)
	if err != nil {
		return models.Store{}, err
	}

	storeData, err := sc.SellerDao.GetStore(store_id)
	if err != nil {
		return models.Store{}, err
	}

	return storeData, nil
}
