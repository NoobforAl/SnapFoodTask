package core

import (
	"Api/db"
	"Api/schema"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func order(c *fiber.Ctx) error {
	var food schema.Food
	if err := c.BodyParser(&food); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"detail": "bad request!",
		})
	}

	data, err := json.Marshal(food)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"detail": fmt.Sprintf("%v", err.Error()),
		})
	}

	db.Redis.LPush("foods", data)
	return c.Status(200).JSON(fiber.Map{
		"detail": "food in possess!",
	})
}

func Notfound(c *fiber.Ctx) error {
	return c.Status(404).JSON(fiber.Map{
		"detail": "Not found!",
	})
}
