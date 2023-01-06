package main

import (
	"Api/core"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{AppName: "Api"})
	app.Use(logger.New(logger.Config{
		Format: "\n${time} | ${ip}:${port} | ${status} - ${method} ${path}",
	})).Name("logger")

	core.Router(app)

	if err := app.Listen("127.0.0.1:8080"); err != nil {
		log.Fatal(err)
	}
}
