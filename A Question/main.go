// wrong: This file only for test
// your dev with air
// this file not compiled

package main

import (
	"A-Question/Api/router"
	"A-Question/NodeFoodAdder/core"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func api() {
	app := fiber.New(fiber.Config{AppName: "Api"})
	app.Use(logger.New(logger.Config{
		Format: "\n${time} | ${ip}:${port} | ${status} - ${method} ${path}",
	})).Name("logger")

	router.Router(app)

	if err := app.Listen("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}

func foodAdder() {
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

func main() {
	go foodAdder()
	go api()

	// this is a trap
	// program wait !
	select {}
}
