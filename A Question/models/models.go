package models

import "gorm.io/gorm"

type FoodSchema struct {
	Order_id uint   `json:"order_id"`
	Price    uint   `json:"price"`
	Title    string `json:"title"`
}

type FoodModel struct {
	gorm.Model
	Order_id uint
	Price    uint
	Title    string
}
