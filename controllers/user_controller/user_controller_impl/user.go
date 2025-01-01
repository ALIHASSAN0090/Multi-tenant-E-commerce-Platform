package user_controller_impl

import (
	"database/sql"
	"ecommerce-platform/Dao"
	"ecommerce-platform/controllers/user_controller"
	"ecommerce-platform/models"
	"ecommerce-platform/utils"
	"fmt"

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
	tx, err := uc.db.Begin()
	if err != nil {
		return models.CreateOrder{}, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	TotalPrice, unitPrices, ItemsTotalPrices, err := uc.UserDao.GetTotalPriceUnitPrice(tx, orderData.OrderItems)
	if err != nil {
		_ = tx.Rollback()
		return models.CreateOrder{}, err
	}

	if TotalPrice == 0 {
		_ = tx.Rollback()
		return models.CreateOrder{}, fmt.Errorf("no items selected or items not found")
	}
	orderData.Order.TotalPrice = TotalPrice

	createdOrder, err := uc.UserDao.CreateOrder(tx, orderData)
	if err != nil {
		_ = tx.Rollback()
		return models.CreateOrder{}, err
	}

	orderData.Order = createdOrder
	for i := range orderData.OrderItems {
		orderData.OrderItems[i].OrderID = createdOrder.ID
		orderData.OrderItems[i].PricePerItem = float64(unitPrices[i])
		orderData.OrderItems[i].TotalPrice = ItemsTotalPrices[i]
	}

	createdItems, err := uc.UserDao.CreateItems(tx, orderData.OrderItems)
	if err != nil {
		_ = tx.Rollback()
		return models.CreateOrder{}, err
	}

	orderData.OrderItems = createdItems

	if err := tx.Commit(); err != nil {
		return models.CreateOrder{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

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
