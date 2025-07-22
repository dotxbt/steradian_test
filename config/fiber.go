package config

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				return ctx.Status(code).JSON(fiber.Map{
					"code":  code,
					"error": err.Error(),
				})
			}
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).SendString("Ups! masih ada kendala di server kami, silahkan mencoba kembali di lain waktu")
			}
			return nil
		},
	}
}
