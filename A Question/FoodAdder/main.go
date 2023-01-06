package main

import (
	"FoodAdder/db"
	"FoodAdder/models"
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Start FoodAdder!")
	for {
		data, err := db.Redis.BLPop(1000*time.Second, "foods").Result()
		if err != nil {
			log.Error(err.Error())
		}

		var food models.FoodPars
		err = json.Unmarshal([]byte(data[1]), &food)
		if err != nil {
			log.Error(err.Error())
		}

		res := db.Sql.Create(&models.Food{
			Order_id: food.Order_id,
			Price:    food.Price,
			Title:    food.Title,
		})

		if res.Error != nil {
			log.Error(res.Error.Error())
			log.Warn("Data Added Agin!")
			db.Redis.LPush("foods", data[1])
		} else {
			log.Println("Food Added!")
		}
	}
}
