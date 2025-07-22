package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/steradian_test/internal/domain/model"
	"github.com/steradian_test/internal/domain/usecase"
)

type CarHandler struct {
	Usecase *usecase.CarUsecase
}

func NewCarHandler(app fiber.Router, usecase *usecase.CarUsecase) {
	handler := &CarHandler{Usecase: usecase}
	app.Get("/cars", handler.GetCars)
	app.Get("/cars/:id", handler.GetCarById)
	app.Post("/cars", handler.CreateCar)
	app.Put("/cars", handler.UpdateCar)
	app.Delete("/cars/:id", handler.DeleteCar)
}

func (h *CarHandler) CreateCar(c *fiber.Ctx) error {
	car := new(model.Car)
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data car",
		})
	}

	car, err := h.Usecase.CreateCar(car)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed create a Car",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Create Car successful",
		"data":    car,
	})
}

func (h *CarHandler) GetCars(c *fiber.Ctx) error {
	cars, err := h.Usecase.GetCars()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ups! something happened in our service. please try again later",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Get Cars successful",
		"data":    cars,
	})
}

func (h *CarHandler) GetCarById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID Id, ID must be a number",
		})
	}
	car, err := h.Usecase.GetCarById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Car Not Found!",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Get Cars successful",
		"data":    car,
	})
}

func (h *CarHandler) UpdateCar(c *fiber.Ctx) error {
	car := new(model.Car)
	if err := c.BodyParser(car); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data car",
		})
	}

	err := h.Usecase.UpdateCar(car)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to update Car",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Update Car successful",
	})
}

func (h *CarHandler) DeleteCar(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID Id, ID must be a number",
		})
	}
	err = h.Usecase.DeleteCar(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Car Not Found!",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Delete Cars successful",
	})
}
