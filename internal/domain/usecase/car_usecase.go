package usecase

import (
	"github.com/steradian_test/internal/domain/model"
	"github.com/steradian_test/internal/domain/repository"
)

type CarUsecase struct {
	repo repository.CarRepository
}

func NewCarUsecase(r repository.CarRepository) *CarUsecase {
	return &CarUsecase{repo: r}
}

func (u *CarUsecase) GetCars() ([]model.Car, error) {
	return u.repo.FindAll()
}

func (u *CarUsecase) GetCarById(id int) (*model.Car, error) {
	return u.repo.FindById(id)
}

func (u *CarUsecase) CreateCar(car *model.Car) (*model.Car, error) {
	return u.repo.Create(car)
}

func (u *CarUsecase) UpdateCar(car *model.Car) (*model.Car, error) {
	return u.repo.Update(car)
}

func (u *CarUsecase) DeleteCar(id int) error {
	return u.repo.Delete(id)
}
