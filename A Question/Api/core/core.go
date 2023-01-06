package core

import "github.com/gofiber/fiber/v2"

func Router(r fiber.Router) {
	r.Post("/api/order", order).Name("Order")
	r.All("*", Notfound).Name("Not Found")
}
