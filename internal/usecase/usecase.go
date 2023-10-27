package usecase

import (
	"WB_Intern_L0/entity"
	"WB_Intern_L0/internal/repo"
)

type Order interface {
	NewOrder(order entity.Order) error
	GetOrderByOrderUID(orderUID string) (entity.Order, error)
}

type Payment interface {
	NewPayment(payment entity.Payment, orderID string) error
	GetPaymentByOrderUID(orderUID string) (entity.Payment, error)
}

type Delivery interface {
	NewDelivery(delivery entity.Delivery, orderID string) error
	GetDeliveryByOrderUID(orderUID string) (entity.Delivery, error)
}

type Items interface {
	NewItem(item entity.Item, orderID string) error
	GetItemsByOrderUID(orderUID string) ([]entity.Item, error)
}

type UseCase struct {
	Order
	Payment
	Delivery
	Items
}

func NewUseCase(repo *repo.Repository) *UseCase {
	return &UseCase{
		Order:    NewOrderUseCase(repo.Order),
		Payment:  NewPaymentUseCase(repo.Payments),
		Delivery: NewDeliveryUseCase(repo.Deliveries),
		Items:    NewItemsUseCase(repo.Items),
	}
}
