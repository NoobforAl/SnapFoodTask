package methods

import (
	"A-Question/db"
	"A-Question/models"
	"encoding/json"
	"time"
)

type Food interface {
	Get() error
	Set() error
}

type FoodSchema models.FoodSchema

// for pars data and set redis database
func (food *FoodSchema) parsToBytes() ([]string, error) {
	data, err := json.Marshal(food)
	return []string{string(data)}, err
}

// get data form redis and pars to FoodSchema
func (food *FoodSchema) parsToFoodSchema(d []byte) error {
	return json.Unmarshal(d, food)
}

// set data food Schema to redis database
func (food *FoodSchema) Get() error {
	data, err := db.Redis.BLPop(1000*time.Second, "foods").Result()
	if err != nil {
		return err
	}

	err = food.parsToFoodSchema([]byte(data[1]))
	if err != nil {
		return err
	}
	return nil
}

// set data food Schema to redis database
func (food *FoodSchema) Set() error {
	// pars food schema to []byte
	f, err := food.parsToBytes()
	if err != nil {
		return err
	}

	re := db.Redis.LPush("foods", f)
	return re.Err()
}

type FoodModel models.FoodModel

// get data from sqlDb
// this only use for test!
func (food *FoodModel) Get() error {
	return db.Sql.First(food, food.ID).Error
}

// set data food in sql database
func (food *FoodModel) Set() error {
	return db.Sql.Create(food).Error
}
