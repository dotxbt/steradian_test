package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/steradian_test/config"
	"github.com/steradian_test/internal/delivery"
)

func main() {
	db := config.InitDB()
	app := fiber.New(config.FiberConfig())
	app.Use(recover.New())
	delivery.CoreDelivery(app, db)
	app.Listen(":3000")
}
