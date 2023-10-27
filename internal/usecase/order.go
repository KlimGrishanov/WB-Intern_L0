package usecase

import (
	"WB_Intern_L0/entity"
	"WB_Intern_L0/internal/repo"
	"errors"
	"go.uber.org/zap"
)

type OrderUseCase struct {
	repo repo.Order
}

func (o OrderUseCase) GetOrders() ([]entity.Order, error) {
	return o.repo.GetOrders()
}

func (o OrderUseCase) NewOrder(order entity.Order) error {
	return o.repo.CreateOrder(order)
}

func (o OrderUseCase) GetOrderByOrderUID(orderUID string) (entity.Order, error) {
	var order entity.Order

	orderArr, err := o.repo.GetOrderByOrderUID(orderUID)

	if err != nil {
		return entity.Order{}, err
	}

	if len(orderArr) != 0 {
		order = orderArr[0]
	} else {
		zap.L().Error("can't find order with this id")
		return entity.Order{}, errors.New("can't find order with this id")
	}

	return order, nil
}

func NewOrderUseCase(repo repo.Order) *OrderUseCase {
	return &OrderUseCase{repo: repo}
}
