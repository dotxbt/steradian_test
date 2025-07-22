package repository

import "github.com/steradian_test/internal/domain/model"

type OrderRepository interface {
	Create(car *model.Order) (*model.Order, error)
	FindAll() ([]model.Order, error)
	FindById(carId int) (*model.Order, error)
	Update(car *model.Order) error
	Delete(carId int) error
}
