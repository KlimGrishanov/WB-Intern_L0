package usecase

import (
	"WB_Intern_L0/entity"
	"WB_Intern_L0/internal/repo"
)

type PaymentUseCase struct {
	repo repo.Payments
}

func (p PaymentUseCase) NewPayment(payment entity.Payment, orderID string) error {
	return p.repo.CreatePayment(payment, orderID)
}

func NewPaymentUseCase(repo repo.Payments) *PaymentUseCase {
	return &PaymentUseCase{repo: repo}
}
