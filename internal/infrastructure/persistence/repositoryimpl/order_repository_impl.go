package repositoryimpl

import (
	"database/sql"
	"fmt"
	"time"

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
	result, err := c.DB.Exec("INSERT INTO orders (car_id, order_date, pickup_date, dropoff_date, pickup_location, dropoff_location) SELECT ?,?,?,?,?,? WHERE NOT EXISTS (SELECT 1 FROM orders WHERE car_id=? AND ((? BETWEEN pickup_date AND dropoff_date) OR (? BETWEEN pickup_date AND dropoff_date))) RETURNING order_id", &order.CarId, order.OrderDate.Format(time.RFC3339), order.PickupDate.Format(time.RFC3339), order.DropoffDate.Format(time.RFC3339), &order.PickupLocation, &order.DropoffLocation, &order.CarId, order.PickupDate.Format(time.RFC3339), order.DropoffDate.Format(time.RFC3339))
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 || id == 0 {
		return nil, fmt.Errorf("RENTED")
	}

	orderId := int(id)
	order.OrderId = &orderId
	return order, nil
}

func (c *OrderRepositoryImp) FindAll() ([]model.Order, error) {
	rows, err := c.DB.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []model.Order
	for rows.Next() {
		var order model.Order
		var orderDateStr, pickupDateStr, dropoffDateStr string

		err = rows.Scan(&order.OrderId, &order.CarId, &orderDateStr, &pickupDateStr, &dropoffDateStr, &order.PickupLocation, &order.DropoffLocation)
		if err != nil {
			//
		}
		order.OrderDate, _ = time.Parse(time.RFC3339, orderDateStr)
		order.PickupDate, _ = time.Parse(time.RFC3339, pickupDateStr)
		order.DropoffDate, _ = time.Parse(time.RFC3339, dropoffDateStr)
		orders = append(orders, order)
	}
	return orders, nil
}

func (c *OrderRepositoryImp) FindById(id int) (*model.Order, error) {
	fmt.Println(id)
	row := c.DB.QueryRow("SELECT * FROM orders WHERE order_id=?", id)
	var order model.Order
	var orderDateStr, pickupDateStr, dropoffDateStr string
	err := row.Scan(&order.OrderId, &order.CarId, &orderDateStr, &pickupDateStr, &dropoffDateStr, &order.PickupLocation, &order.DropoffLocation)
	if err != nil {
		return nil, err
	}
	order.OrderDate, _ = time.Parse(time.RFC3339, orderDateStr)
	order.PickupDate, _ = time.Parse(time.RFC3339, pickupDateStr)
	order.DropoffDate, _ = time.Parse(time.RFC3339, dropoffDateStr)
	return &order, nil
}

func (c *OrderRepositoryImp) Update(order *model.Order) error {
	_, err := c.DB.Exec("UPDATE orders SET order_date=?, pickup_date=?, dropoff_date=?, pickup_location=?, dropoff_location=? WHERE order_id=?", order.OrderDate.Format(time.RFC3339), order.PickupDate.Format(time.RFC3339), order.DropoffDate.Format(time.RFC3339), order.PickupLocation, &order.DropoffLocation, &order.OrderId)
	if err != nil {
		return err
	}
	return nil
}

func (c *OrderRepositoryImp) Delete(orderId int) error {
	_, err := c.DB.Exec("DELETE FROM orders WHERE order_id=?", orderId)
	if err != nil {
		return err
	}
	return nil
}
