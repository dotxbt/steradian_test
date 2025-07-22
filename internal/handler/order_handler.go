package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/steradian_test/internal/domain/model"
	"github.com/steradian_test/internal/domain/usecase"
)

type OrderHandler struct {
	Usecase *usecase.OrderUsecase
}

func NewOrderHandler(app fiber.Router, usecase *usecase.OrderUsecase) {
	handler := &OrderHandler{Usecase: usecase}
	app.Get("/orders", handler.GetOrders)
	app.Get("/orders/:id", handler.GetOrderById)
	app.Post("/orders", handler.CreateOrder)
	app.Put("/orders", handler.UpdateOrder)
	app.Delete("/orders/:id", handler.DeleteOrder)
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	Order := new(model.Order)
	if err := c.BodyParser(Order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data Order Request",
		})
	}

	Order, err := h.Usecase.CreateOrder(Order)
	if err != nil {
		errs := err.Error()
		if errs == "RENTED" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Car is being rented",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed create a Order" + err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Create Order successful",
		"data":    Order,
	})
}

func (h *OrderHandler) GetOrders(c *fiber.Ctx) error {
	Orders, err := h.Usecase.GetOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Ups! something happened in our service. please try again later",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Get Orders successful",
		"data":    Orders,
	})
}

func (h *OrderHandler) GetOrderById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID Id, ID must be a number",
		})
	}
	Order, err := h.Usecase.GetOrderById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Order Not Found!",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Get Orders successful",
		"data":    Order,
	})
}

func (h *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	Order := new(model.Order)
	if err := c.BodyParser(Order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data Order",
		})
	}

	err := h.Usecase.UpdateOrder(Order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to update Order",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Update Order successful",
	})
}

func (h *OrderHandler) DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID Id, ID must be a number",
		})
	}
	err = h.Usecase.DeleteOrder(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Order Not Found!",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Delete Orders successful",
	})
}
