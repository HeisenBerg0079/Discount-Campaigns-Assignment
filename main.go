package main

import (
	"github.com/Earth/discount-campaigns/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/calculate-discount", handler.CalculateDiscountedPrice)

	app.Listen(":8080")
}
