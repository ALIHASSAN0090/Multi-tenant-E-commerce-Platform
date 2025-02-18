package Dao

import (
	"database/sql"
	"ecommerce-platform/models"
	"ecommerce-platform/utils"
	"fmt"
	"time"
)

func NewUserDao(db *sql.DB) UserDao {
	return &UserDaoImpl{
		db: db,
	}
}

type UserDaoImpl struct {
	db *sql.DB
}

func (dao *UserDaoImpl) CreateItems(tx *sql.Tx, items []models.OrderItem) ([]models.OrderItem, error) {
	var createdItems []models.OrderItem
	query := `INSERT INTO order_items (order_id, item_id, quantity, price_per_item, total_price, created_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, order_id, item_id, quantity, price_per_item, total_price`

	for _, item := range items {
		var createdItem models.OrderItem
		var err error
		if tx != nil {
			err = tx.QueryRow(query, item.OrderID, item.ID, item.Quantity, item.PricePerItem, item.TotalPrice, time.Now()).
				Scan(&createdItem.ID, &createdItem.OrderID, &createdItem.ItemID, &createdItem.Quantity, &createdItem.PricePerItem, &createdItem.TotalPrice)
		} else {
			err = dao.db.QueryRow(query, item.OrderID, item.ID, item.Quantity, item.PricePerItem, item.TotalPrice, time.Now()).
				Scan(&createdItem.ID, &createdItem.OrderID, &createdItem.ItemID, &createdItem.Quantity, &createdItem.PricePerItem, &createdItem.TotalPrice)
		}
		if err != nil {
			return nil, err
		}
		createdItems = append(createdItems, createdItem)
	}

	return createdItems, nil
}

func (dao *UserDaoImpl) GetTotalPriceUnitPrice(tx *sql.Tx, items []models.OrderItem) (float64, []int64, []float64, error) {
	var totalPrice float64
	var unitPrices []int64
	var itemsPrices []float64

	for _, item := range items {
		var price, finalPrice float64
		var discount int64
		query := `SELECT price, discount FROM items WHERE id = $1`

		var err error
		if tx != nil {
			err = tx.QueryRow(query, item.ID).Scan(&price, &discount)
		} else {
			err = dao.db.QueryRow(query, item.ID).Scan(&price, &discount)
		}

		if err != nil {
			return 0, nil, nil, err
		}

		if discount > 0 {
			discountedPrice, err := utils.GetDiscountedPrice(float32(price), discount)
			if err != nil {
				return 0, nil, nil, err
			}
			finalPrice = discountedPrice * float64(item.Quantity)
		} else {
			finalPrice = price * float64(item.Quantity)
		}

		itemsPrices = append(itemsPrices, finalPrice)
		totalPrice += finalPrice
		unitPrices = append(unitPrices, int64(price))
	}

	return totalPrice, unitPrices, itemsPrices, nil
}

func (dao *UserDaoImpl) CreateOrder(tx *sql.Tx, orderData models.CreateOrder) (models.Order, error) {
	query := `INSERT INTO orders(user_id, store_id, total_price, status, created_at, created_by)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, user_id, store_id, total_price, status, created_at, created_by`

	var order models.Order
	var err error
	if tx != nil {
		err = tx.QueryRow(query, orderData.Order.UserID, orderData.Order.StoreID, orderData.Order.TotalPrice, "pending", time.Now(), orderData.Order.CreatedBy).
			Scan(&order.ID, &order.UserID, &order.StoreID, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.CreatedBy)
	} else {
		err = dao.db.QueryRow(query, orderData.Order.UserID, orderData.Order.StoreID, orderData.Order.TotalPrice, "pending", time.Now(), orderData.Order.CreatedBy).
			Scan(&order.ID, &order.UserID, &order.StoreID, &order.TotalPrice, &order.Status, &order.CreatedAt, &order.CreatedBy)
	}

	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (dao *UserDaoImpl) CreateSeller(seller models.SellerStore) (models.Seller, error) {

	exists, err := dao.checkExistingSeller(seller.UserID)
	if err != nil {
		return models.Seller{}, err
	}
	if exists {
		return models.Seller{}, fmt.Errorf("seller with user_id %d already exists", seller.UserID)
	}

	query := `INSERT INTO seller(user_id, business_name, contact_number, created_at, active, seller_img, cnic_number, cnic_image) 
	values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	var sellerID int64
	err = dao.db.QueryRow(query, seller.UserID, seller.BusinessName, seller.ContactNumber, time.Now(), true, seller.SellerImg, seller.CNICNumber, seller.CNICImage).Scan(&sellerID)
	if err != nil {
		return models.Seller{}, err
	}

	return models.Seller{
		ID:            sellerID,
		UserID:        seller.UserID,
		BusinessName:  seller.BusinessName,
		ContactNumber: seller.ContactNumber,
		CreatedAt:     time.Now(),
		Active:        seller.Active,
		SellerImg:     seller.SellerImg,
		CNICNumber:    seller.CNICNumber,
		CNICImage:     seller.CNICImage,
	}, nil
}

func (dao *UserDaoImpl) CreateStore(store models.SellerStore, seller_id int64) (models.Store, error) {
	query := `INSERT INTO stores(seller_id, store_img, store_name, store_address, store_description, created_at) 
	values ($1, $2, $3, $4, $5, $6) returning id`

	var storeID int64
	err := dao.db.QueryRow(query, seller_id, store.StoreImg, store.StoreName, store.StoreAddress, store.StoreDescription, time.Now()).Scan(&storeID)
	if err != nil {
		return models.Store{}, err
	}

	return models.Store{
		ID:               storeID,
		SellerID:         seller_id,
		StoreImg:         store.StoreImg,
		StoreName:        store.StoreName,
		StoreAddress:     store.StoreAddress,
		StoreDescription: store.StoreDescription,
		CreatedAt:        time.Now(),
	}, nil
}

func (dao *UserDaoImpl) checkExistingSeller(userID int64) (bool, error) {
	existingSellerQuery := `SELECT id FROM seller WHERE user_id = $1`
	var existingSellerID int64
	err := dao.db.QueryRow(existingSellerQuery, userID).Scan(&existingSellerID)
	if err == nil {
		return true, fmt.Errorf("user_id already exists in the seller table")
	} else if err != sql.ErrNoRows {
		return false, err
	}
	return false, nil
}

func (dao *UserDaoImpl) ChangeRoleToSeller(id int64) (bool, error) {

	var roleID int
	query2 := `SELECT id FROM roles WHERE name = 'seller'`
	err := dao.db.QueryRow(query2).Scan(&roleID)
	if err != nil {
		return false, err
	}

	query := `UPDATE users SET role_id = $1 WHERE id = $2`

	_, err = dao.db.Exec(query, roleID, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (dao *UserDaoImpl) GetStores() ([]models.Store, error) {
	query := `SELECT id, store_img, store_name, store_description FROM stores`
	rows, err := dao.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []models.Store
	for rows.Next() {
		var store models.Store
		if err := rows.Scan(&store.ID, &store.StoreImg, &store.StoreName, &store.StoreDescription); err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stores, nil
}

func (dao *UserDaoImpl) GetStoreItems(store_id int64) (models.StoreItems, error) {
	query := `SELECT s.id, s.store_img, s.store_name, s.store_description, i.id, i.name, i.description, i.price, i.discount 
	          FROM stores AS s
	          JOIN items AS i ON i.store_id = s.id
	          WHERE s.id = $1`

	rows, err := dao.db.Query(query, store_id)
	if err != nil {
		return models.StoreItems{}, err
	}
	defer rows.Close()

	var storeItems models.StoreItems
	var initialized bool

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&storeItems.ID, &storeItems.StoreImg, &storeItems.StoreName, &storeItems.StoreDescription,
			&item.ID, &item.Name, &item.Description, &item.Price, &item.Discount); err != nil {
			return models.StoreItems{}, err
		}

		if !initialized {
			storeItems = models.StoreItems{
				ID:       storeItems.ID,
				StoreImg: storeItems.StoreImg,

				StoreName:        storeItems.StoreName,
				StoreDescription: storeItems.StoreDescription,
				Items:            []models.Item{},
			}
			initialized = true
		}

		storeItems.Items = append(storeItems.Items, item)
	}

	if err := rows.Err(); err != nil {
		return models.StoreItems{}, err
	}

	return storeItems, nil
}
