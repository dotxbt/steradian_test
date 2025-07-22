package main

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/steradian_test/config"
	"github.com/steradian_test/internal/delivery"
)

func main() {
	db := config.InitDB()
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				return ctx.Status(code).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Ups! masih ada kendala di server kami, silahkan mencoba kembali di lain waktu")
			}

			// Return from handler
			return nil
		},
	})
	app.Use(recover.New())
	delivery.CoreDelivery(app, db)
	app.Listen(":3000")
}
