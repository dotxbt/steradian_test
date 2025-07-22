package repository

import "github.com/steradian_test/internal/domain/model"

type CarRepository interface {
	Create(car *model.Car) (*model.Car, error)
	FindAll() ([]model.Car, error)
	FindById(carId int) (*model.Car, error)
	Update(car *model.Car) error
	Delete(carId int) error
}
