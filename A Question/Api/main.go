package main

import (
	"A-Question/Api/router"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "Api"})
	app.Use(logger.New(logger.Config{
		Format: "\n${time} | ${ip}:${port} | ${status} - ${method} ${path}",
	})).Name("logger")

	router.Router(app)

	if err := app.Listen("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
