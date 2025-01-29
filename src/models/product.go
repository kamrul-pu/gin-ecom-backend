package models

import "gorm.io/gorm"

type Product struct{
	gorm.Model
	Name string `json:"name"`
	Catrgory string `json:"category"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Stock int `json:"stock"`
}