package usecase

import (
	"WB_Intern_L0/entity"
	"WB_Intern_L0/internal/repo"
)

type OrderUseCase struct {
	repo repo.Order
}

func (o OrderUseCase) NewOrder(order entity.Order) error {
	return o.repo.CreateOrder(order)
}

func NewOrderUseCase(repo repo.Order) *OrderUseCase {
	return &OrderUseCase{repo: repo}
}
