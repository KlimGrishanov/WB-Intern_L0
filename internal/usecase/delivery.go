package usecase

import (
	"WB_Intern_L0/entity"
	"WB_Intern_L0/internal/repo"
)

type DeliveryUseCase struct {
	repo repo.Deliveries
}

func (p DeliveryUseCase) NewDelivery(delivery entity.Delivery, orderID string) error {
	return p.repo.CreateDelivery(delivery, orderID)
}

func NewDeliveryUseCase(repo repo.Deliveries) *DeliveryUseCase {
	return &DeliveryUseCase{repo: repo}
}
