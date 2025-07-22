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
	if order.PickupDate.After(order.DropoffDate) {
		return nil, fmt.Errorf("tanggal pickup harus sebelum dropoff")
	}

	row := c.DB.QueryRow("INSERT INTO orders (car_id, pickup_date, dropoff_date, pickup_location, dropoff_location) SELECT ?,?,?,?,? WHERE NOT EXISTS (SELECT 1 FROM orders WHERE car_id=? AND pickup_date <= ?  AND dropoff_date >= ?) RETURNING *", &order.CarId, order.PickupDate.Format(time.RFC3339), order.DropoffDate.Format(time.RFC3339), &order.PickupLocation, &order.DropoffLocation, &order.CarId, order.DropoffDate.Format(time.RFC3339), order.PickupDate.Format(time.RFC3339))
	var newOrder model.Order
	var orderDateStr, pickupDateStr, dropoffDateStr string
	err := row.Scan(&newOrder.OrderId, &newOrder.CarId, &orderDateStr, &pickupDateStr, &dropoffDateStr, &newOrder.PickupLocation, &newOrder.DropoffLocation)

	if err != nil {
		return nil, fmt.Errorf("RENTED")
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
	var orders []model.Order
	for rows.Next() {
		var order model.Order
		var orderDateStr, pickupDateStr, dropoffDateStr string

		err = rows.Scan(&order.OrderId, &order.CarId, &orderDateStr, &pickupDateStr, &dropoffDateStr, &order.PickupLocation, &order.DropoffLocation)
		if err != nil {
			//
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
	fmt.Println(id)
	row := c.DB.QueryRow("SELECT * FROM orders WHERE order_id=?", id)
	var order model.Order
	var orderDateStr, pickupDateStr, dropoffDateStr string
	err := row.Scan(&order.OrderId, &order.CarId, &orderDateStr, &pickupDateStr, &dropoffDateStr, &order.PickupLocation, &order.DropoffLocation)
	if err != nil {
		return nil, err
	}
	newOd, _ := time.Parse(time.RFC3339, orderDateStr)
	order.OrderDate = &newOd
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
