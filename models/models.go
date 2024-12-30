package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Id    uint    `json:"id"`
	Email *string `json:"email"`
	Name  *string `json:"name"`
	Role  *string `json:"role"`
	jwt.StandardClaims
}

type User struct {
	ID           int64      `json:"id" db:"id"`
	ProfileImg   string     `json:"profile_img" db:"profile_img"`
	RoleID       int64      `json:"role_id" db:"role_id"`
	Name         string     `json:"name" db:"name"`
	Email        string     `json:"email" db:"email"`
	HashPassword string     `json:"hash_password" db:"hash_password"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type SellerStore struct {
	ID               int64      `json:"id" db:"id"`
	SellerImg        string     `json:"seller_img" db:"seller_img"`
	CNICNumber       string     `json:"cnic_number" db:"cnic_number"`
	CNICImage        string     `json:"cnic_image" db:"cnic_image"`
	UserID           int64      `json:"user_id" db:"user_id"`
	Active           bool       `json:"active" db:"active" default:"true"`
	BusinessName     string     `json:"business_name" db:"business_name"`
	ContactNumber    string     `json:"contact_number" db:"contact_number"`
	StoreImg         string     `json:"store_img" db:"store_img"`
	SellerID         int64      `json:"seller_id" db:"seller_id"`
	StoreName        string     `json:"store_name" db:"store_name"`
	StoreDescription string     `json:"store_description" db:"store_description"`
	StoreAddress     string     `json:"store_address" db:"store_address"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Store struct {
	ID               int64      `json:"id" db:"id"`
	StoreImg         string     `json:"store_img" db:"store_img"`
	SellerID         int64      `json:"seller_id" db:"seller_id"`
	StoreName        string     `json:"store_name" db:"store_name"`
	StoreDescription string     `json:"store_description" db:"store_description"`
	StoreAddress     string     `json:"store_address" db:"store_address"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Role struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Seller model
type Seller struct {
	ID            int64      `json:"id" db:"id"`
	Active        bool       `json:"active" db:"active" default:"true"`
	SellerImg     string     `json:"seller_img" db:"seller_img"`
	CNICNumber    string     `json:"cnic_number" db:"cnic_number"`
	CNICImage     string     `json:"cnic_image" db:"cnic_image"`
	UserID        int64      `json:"user_id" db:"user_id"`
	BusinessName  string     `json:"business_name" db:"business_name"`
	ContactNumber string     `json:"contact_number" db:"contact_number"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type Item struct {
	ID              int64      `json:"id" db:"id"`
	ItemImg         string     `json:"item_img" db:"item_img"`
	StoreID         int64      `json:"store_id" db:"store_id"`
	Name            string     `json:"name" db:"name"`
	Description     string     `json:"description" db:"description"`
	Price           float64    `json:"price" db:"price"`
	DiscountedPrice float64    `json:"discounted_price,omitempty" db:"discounted_price"`
	StockQuantity   int        `json:"stock_quantity" db:"stock_quantity"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	Discount        int64      `json:"discount,omitempty" db:"discount"`
}

type StoreItems struct {
	ID               int64  `json:"id" db:"id"`
	StoreImg         string `json:"store_img" db:"store_img"`
	SellerID         int64  `json:"seller_id" db:"seller_id"`
	StoreName        string `json:"store_name" db:"store_name"`
	StoreDescription string `json:"store_description" db:"store_description"`
	StoreAddress     string `json:"store_address" db:"store_address"`
	Items            []Item `json:"items" db:"items"`
}

// Order model
type Order struct {
	ID         int64      `json:"id" db:"id"`
	UserID     int64      `json:"user_id" db:"user_id"`
	StoreID    int64      `json:"store_id" db:"store_id"`
	TotalPrice float64    `json:"total_price" db:"total_price"`
	Status     string     `json:"status" db:"status"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	CreatedBy  int64      `json:"created_by" db:"created_by"`
	UpdatedBy  *int64     `json:"updated_by,omitempty" db:"updated_by"`
}

type OrderItem struct {
	ID           int64      `json:"id" db:"id"`
	OrderID      int64      `json:"order_id" db:"order_id"`
	ItemID       int64      `json:"item_id" db:"item_id"`
	Quantity     int        `json:"quantity" db:"quantity"`
	PricePerItem float64    `json:"price_per_item" db:"price_per_item"`
	TotalPrice   float64    `json:"total_price" db:"total_price"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type CreateOrder struct {
	Order      Order       `json:"order" db:"order"`
	OrderItems []OrderItem `json:"order_items" db:"order_items"`
}

// Payment model
type Payment struct {
	ID            int64      `json:"id" db:"id"`
	OrderID       int64      `json:"order_id" db:"order_id"`
	PaymentMethod string     `json:"payment_method" db:"payment_method"`
	PaymentStatus string     `json:"payment_status" db:"payment_status"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
	CreatedBy     int64      `json:"created_by" db:"created_by"`
	UpdatedBy     *int64     `json:"updated_by,omitempty" db:"updated_by"`
}

// OrderHistory model
type OrderHistory struct {
	ID        int64      `json:"id" db:"id"`
	OrderID   int64      `json:"order_id" db:"order_id"`
	Status    string     `json:"status" db:"status"`
	ChangedAt *time.Time `json:"changed_at,omitempty" db:"changed_at"`
	ChangedBy int64      `json:"changed_by" db:"changed_by"`
}

// Review model
type Review struct {
	ID        int64     `json:"id" db:"id"`
	OrderID   int64     `json:"order_id" db:"order_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Rating    int       `json:"rating" db:"rating"`
	Comment   string    `json:"comment" db:"comment"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
