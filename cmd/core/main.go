package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/steradian_test/config"
	"github.com/steradian_test/internal/delivery"
)

func main() {
	db := config.InitDB()
	app := fiber.New()
	delivery.CoreDelivery(app, db)
	app.Listen(":3000")
}
