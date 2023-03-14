package core

import (
	"A-Question/methods"
	"fmt"
)

func GetFood() (string, error) {
	// get data first food in list
	var ffood methods.Food = &methods.FoodSchema{}
	if err := ffood.Get(); err != nil {
		return "", err
	}

	// setup all value from ffood
	var food methods.Food = &methods.FoodModel{}
	food.(*methods.FoodModel).Order_id = ffood.(*methods.FoodSchema).Order_id
	food.(*methods.FoodModel).Price = ffood.(*methods.FoodSchema).Price
	food.(*methods.FoodModel).Title = ffood.(*methods.FoodSchema).Title

	// set data on database
	if err := food.Set(); err != nil {
		return "", err
	}

	return fmt.Sprintf("This Food Add Title: %s", ffood.(*methods.FoodSchema).Title), nil
}
