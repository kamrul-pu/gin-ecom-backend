package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID           uint            `json:"user_id"`
	ShippingAddressID uint           `json:"shipping_address_id"`
	ShippingAddress  ShippingAddress `json:"shipping_address" gorm:"foreignKey:ShippingAddressID"`
	OrderItems       []OrderItem     `json:"order_items"`
	Total            float64         `json:"total"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type ShippingAddress struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
}