package main

import (
	"A-Question/NodeFoodAdder/core"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Start FoodAdder!")
	for {
		msg, err := core.GetFood()
		if err != nil {
			log.Error(err)
		} else {
			log.Info(msg)
		}
	}
}
