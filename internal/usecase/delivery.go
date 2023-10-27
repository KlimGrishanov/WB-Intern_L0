package usecase

import (
	"WB_Intern_L0/entity"
	"WB_Intern_L0/internal/repo"
	"errors"
	"go.uber.org/zap"
)

type DeliveryUseCase struct {
	repo repo.Deliveries
}

func (p DeliveryUseCase) GetDeliveryByOrderUID(orderUID string) (entity.Delivery, error) {
	var delivery entity.Delivery

	deliveryArr, err := p.repo.GetDeliveryByOrderUID(orderUID)

	if err != nil {
		return entity.Delivery{}, err
	}

	if len(deliveryArr) != 0 {
		delivery = deliveryArr[0]
	} else {
		zap.L().Error("can't find order with this id")
		return entity.Delivery{}, errors.New("can't find order with this id")
	}

	return delivery, nil
}

func (p DeliveryUseCase) NewDelivery(delivery entity.Delivery, orderID string) error {
	return p.repo.CreateDelivery(delivery, orderID)
}

func NewDeliveryUseCase(repo repo.Deliveries) *DeliveryUseCase {
	return &DeliveryUseCase{repo: repo}
}
