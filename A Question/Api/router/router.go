package router

import (
	"A-Question/Api/controller"

	"github.com/gofiber/fiber/v2"
)

func Router(r fiber.Router) {
	r.Post("/api/order", controller.Order).Name("Order")
	r.All("*", controller.Notfound).Name("Not Found")
}
