package usecase

import (
	"github.com/steradian_test/internal/domain/model"
	"github.com/steradian_test/internal/domain/repository"
)

type OrderUsecase struct {
	repo repository.OrderRepository
}

func NewOrderUsecase(r repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: r}
}

func (u *OrderUsecase) GetOrders() ([]model.Order, error) {
	return u.repo.FindAll()
}

func (u *OrderUsecase) GetOrderById(id int) (*model.Order, error) {
	return u.repo.FindById(id)
}

func (u *OrderUsecase) CreateOrder(Order *model.Order) (*model.Order, error) {
	return u.repo.Create(Order)
}

func (u *OrderUsecase) UpdateOrder(Order *model.Order) error {
	return u.repo.Update(Order)
}

func (u *OrderUsecase) DeleteOrder(id int) error {
	return u.repo.Delete(id)
}
