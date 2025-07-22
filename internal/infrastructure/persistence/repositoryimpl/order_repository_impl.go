package repositoryimpl

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/steradian_test/internal/domain/model"
)

type OrderRepositoryImp struct {
	DB *sql.DB
}

func NewOrderRepositoryImpl(db *sql.DB) *OrderRepositoryImp {
	return &OrderRepositoryImp{
		DB: db,
	}
}

func (c *OrderRepositoryImp) Create(order *model.Order) (*model.Order, error) {
	if order.PickupDate.After(order.DropoffDate) {
		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			"Pickup Date must be lower than or equals Dropoff Date")
	}

	query := `
	INSERT INTO orders (car_id, pickup_date, dropoff_date, pickup_location, dropoff_location) 
	SELECT ?,?,?,?,? WHERE NOT EXISTS 
	(SELECT 1 FROM orders WHERE car_id=? AND pickup_date <= ?  AND dropoff_date >= ?) 
	RETURNING *
	`
	row := c.DB.QueryRow(
		query,
		&order.CarId,
		order.PickupDate.Format(time.RFC3339),
		order.DropoffDate.Format(time.RFC3339),
		&order.PickupLocation,
		&order.DropoffLocation,
		&order.CarId,
		order.DropoffDate.Format(time.RFC3339),
		order.PickupDate.Format(time.RFC3339))

	var newOrder model.Order
	var orderDateStr, pickupDateStr, dropoffDateStr string

	err := row.Scan(
		&newOrder.OrderId,
		&newOrder.CarId,
		&orderDateStr,
		&pickupDateStr,
		&dropoffDateStr,
		&newOrder.PickupLocation,
		&newOrder.DropoffLocation)

	if err != nil {
		msg := "Ups! something error!"
		if strings.Contains(err.Error(), "FOREIGN KEY") {
			msg = "Car not found!"
		}

		if strings.Contains(err.Error(), "no rows in result set") {
			msg = "Car is being rented"
		}

		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			msg,
		)
	}
	newOd, _ := time.Parse(time.RFC3339, orderDateStr)
	newOrder.OrderDate = &newOd
	newOrder.PickupDate, _ = time.Parse(time.RFC3339, pickupDateStr)
	newOrder.DropoffDate, _ = time.Parse(time.RFC3339, dropoffDateStr)
	return &newOrder, nil
}

func (c *OrderRepositoryImp) FindAll() ([]model.Order, error) {
	rows, err := c.DB.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []model.Order{}
	for rows.Next() {
		var order model.Order
		var orderDateStr, pickupDateStr, dropoffDateStr string

		err = rows.Scan(
			&order.OrderId,
			&order.CarId,
			&orderDateStr,
			&pickupDateStr,
			&dropoffDateStr,
			&order.PickupLocation,
			&order.DropoffLocation)

		if err != nil {
			// error to scan
		}
		newOd, _ := time.Parse(time.RFC3339, orderDateStr)
		order.OrderDate = &newOd
		order.PickupDate, _ = time.Parse(time.RFC3339, pickupDateStr)
		order.DropoffDate, _ = time.Parse(time.RFC3339, dropoffDateStr)
		orders = append(orders, order)
	}
	return orders, nil
}

func (c *OrderRepositoryImp) FindById(id int) (*model.Order, error) {
	row := c.DB.QueryRow("SELECT * FROM orders WHERE order_id=?", id)
	var order model.Order
	var orderDateStr, pickupDateStr, dropoffDateStr string

	err := row.Scan(
		&order.OrderId,
		&order.CarId,
		&orderDateStr,
		&pickupDateStr,
		&dropoffDateStr,
		&order.PickupLocation,
		&order.DropoffLocation)

	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			"Order not found!",
		)
	}
	newOd, _ := time.Parse(time.RFC3339, orderDateStr)
	order.OrderDate = &newOd
	order.PickupDate, _ = time.Parse(time.RFC3339, pickupDateStr)
	order.DropoffDate, _ = time.Parse(time.RFC3339, dropoffDateStr)
	return &order, nil
}

func (c *OrderRepositoryImp) Update(order *model.Order) (*model.Order, error) {
	query := `
	UPDATE orders SET order_date=?, pickup_date=?, dropoff_date=?, pickup_location=?, dropoff_location=? 
	WHERE order_id=? RETURNING *
	`
	row := c.DB.QueryRow(
		query,
		order.OrderDate.Format(time.RFC3339),
		order.PickupDate.Format(time.RFC3339),
		order.DropoffDate.Format(time.RFC3339),
		order.PickupLocation,
		&order.DropoffLocation,
		&order.OrderId)

	var updatedOrder model.Order
	var orderDateStr, pickupDateStr, dropoffDateStr string
	err := row.Scan(
		&updatedOrder.OrderId,
		&updatedOrder.CarId,
		&orderDateStr,
		&pickupDateStr,
		&dropoffDateStr,
		&updatedOrder.PickupLocation,
		&updatedOrder.DropoffLocation)

	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			"Failed to update order",
		)
	}

	newOd, _ := time.Parse(time.RFC3339, orderDateStr)
	updatedOrder.OrderDate = &newOd
	updatedOrder.PickupDate, _ = time.Parse(time.RFC3339, pickupDateStr)
	updatedOrder.DropoffDate, _ = time.Parse(time.RFC3339, dropoffDateStr)
	return &updatedOrder, nil
}

func (c *OrderRepositoryImp) Delete(orderId int) (*string, error) {
	var deletedId int
	err := c.DB.QueryRow("DELETE FROM orders WHERE order_id=? RETURNING order_id", orderId).Scan(&deletedId)
	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			"Order not found!",
		)
	}
	res := "Delete order with ID " + fmt.Sprint(deletedId) + " successfull"
	return &res, nil
}
