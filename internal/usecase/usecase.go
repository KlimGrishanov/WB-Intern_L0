package usecase

import (
	"WB_Intern_L0/entity"
	"WB_Intern_L0/internal/repo"
)

type Order interface {
	NewOrder(order entity.Order) error
}

type Payment interface {
	NewPayment(payment entity.Payment, orderID string) error
}

type Delivery interface {
	NewDelivery(delivery entity.Delivery, orderID string) error
}

type Items interface {
	NewItem(delivery entity.Item, orderID string) error
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
