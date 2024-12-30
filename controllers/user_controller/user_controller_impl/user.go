package user_controller_impl

import (
	"database/sql"
	"ecommerce-platform/Dao"
	"ecommerce-platform/controllers/user_controller"
	"ecommerce-platform/models"
	"ecommerce-platform/utils"

	"github.com/gin-gonic/gin"
)

type UsercontrollerImpl struct {
	UserDao Dao.UserDao
	db      *sql.DB
}

type UserControllerConfig struct {
	UserDao Dao.UserDao
	DB      *sql.DB
}

func NewUserImpl(config UserControllerConfig) user_controller.UserControllerConfig {
	return &UsercontrollerImpl{
		UserDao: config.UserDao,
		db:      config.DB,
	}
}

func (uc *UsercontrollerImpl) CreateOrder(c *gin.Context, orderData models.CreateOrder) (models.CreateOrder, error) {
	TotalPrice, err := uc.UserDao.GetTotalPriceUnitPrice(orderData.OrderItems)
	if err != nil {
		utils.HandleError(err)
		return models.CreateOrder{}, err
	}

	orderData.Order.TotalPrice = TotalPrice

	orderId, err := uc.UserDao.CreateOrder(orderData)
	if err != nil {
		utils.HandleError(err)
		return models.CreateOrder{}, err
	}

	createdItems, err := uc.UserDao.CreateItems(orderId, orderData.OrderItems)
	if err != nil {
		utils.HandleError(err)
		return models.CreateOrder{}, err
	}

	orderData.OrderItems = createdItems

	return orderData, nil
}
func (uc *UsercontrollerImpl) CreateSellerStore(c *gin.Context, seller models.SellerStore) (models.Seller, models.Store, error) {
	sellerData, err := uc.UserDao.CreateSeller(seller)
	if err != nil {
		utils.HandleError(err)
		return models.Seller{}, models.Store{}, err
	}

	storeData, err := uc.UserDao.CreateStore(seller, sellerData.ID)
	if err != nil {
		utils.HandleError(err)
		return models.Seller{}, models.Store{}, err
	}

	changed, err := uc.UserDao.ChangeRoleToSeller(sellerData.UserID)
	if err != nil {
		utils.HandleError(err)
		return models.Seller{}, models.Store{}, err
	}

	if !changed {
		return models.Seller{}, models.Store{}, err
	}

	return sellerData, storeData, nil
}

func (uc *UsercontrollerImpl) GetStores(c *gin.Context) ([]models.Store, error) {
	return uc.UserDao.GetStores()
}

func (uc *UsercontrollerImpl) GetStoreItems(c *gin.Context, store_id int64) (models.StoreItems, error) {

	return uc.UserDao.GetStoreItems(store_id)
}
