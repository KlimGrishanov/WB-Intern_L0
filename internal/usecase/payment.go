package usecase

import (
	"WB_Intern_L0/entity"
	"WB_Intern_L0/internal/repo"
	"errors"
	"go.uber.org/zap"
)

type PaymentUseCase struct {
	repo repo.Payments
}

func (p PaymentUseCase) GetPaymentByOrderUID(orderUID string) (entity.Payment, error) {
	var payment entity.Payment

	paymentArr, err := p.repo.GetPaymentsByOrderUID(orderUID)

	if err != nil {
		return entity.Payment{}, err
	}

	if len(paymentArr) != 0 {
		payment = paymentArr[0]
	} else {
		zap.L().Error("can't find order with this id")
		return entity.Payment{}, errors.New("can't find order with this id")
	}

	return payment, nil
}

func (p PaymentUseCase) NewPayment(payment entity.Payment, orderID string) error {
	return p.repo.CreatePayment(payment, orderID)
}

func NewPaymentUseCase(repo repo.Payments) *PaymentUseCase {
	return &PaymentUseCase{repo: repo}
}
