package controller

import (
	"A-Question/methods"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Order(c *fiber.Ctx) error {
	var food methods.Food = &methods.FoodSchema{}

	if err := c.BodyParser(&food); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"detail": "bad request!",
		})
	}

	if err := food.Set(); err != nil {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"detail": "food in possess!",
	})
}

// handel all Unknown request
func Notfound(c *fiber.Ctx) error {
	return c.Status(http.StatusNotFound).JSON(fiber.Map{
		"detail": "Not found!",
	})
}
