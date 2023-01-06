package models

import "gorm.io/gorm"

type FoodPars struct {
	Order_id uint   `json:"order_id"`
	Price    uint   `json:"price"`
	Title    string `json:"title"`
}

type Food struct {
	gorm.Model
	Order_id uint
	Price    uint
	Title    string
}
