package repo

import (
	"WB_Intern_L0/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PaymentsPostgres struct {
	db *sqlx.DB
}

func (p PaymentsPostgres) CreatePayment(payment entity.Payment, orderUID string) error {
	query := fmt.Sprintf("INSERT INTO %s (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)"+" values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", paymentsTable)
	_, err := p.db.Exec(query, orderUID, payment.RequestId, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	return err
}

func (p PaymentsPostgres) GetPaymentsByOrderUID(orderUID string) ([]entity.Payment, error) {
	var payments []entity.Payment

	query := fmt.Sprintf("SELECT * FROM %s WHERE transaction=$1", paymentsTable)
	err := p.db.Select(&payments, query, orderUID)

	return payments, err
}

func (p PaymentsPostgres) DeletePayment(orderUID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE order_uid=$1", paymentsTable)
	_, err := p.db.Exec(query, orderUID)
	return err
}

func NewPaymentsPostgres(db *sqlx.DB) *PaymentsPostgres {
	return &PaymentsPostgres{db: db}
}
