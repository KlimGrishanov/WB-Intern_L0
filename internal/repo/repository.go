package repo

import (
	"WB_Intern_L0/entity"
	"github.com/jmoiron/sqlx"
)

const (
	ordersTable     = "orders"
	deliveriesTable = "deliveries"
	paymentsTable   = "payments"
	itemsTable      = "items"
)

type Items interface {
	CreateItem(item entity.Item, orderUID string) error
	DeleteItem(itemID int) error
	GetItemsByOrderUID(orderUID string) ([]entity.Item, error)
}

type Payments interface {
	CreatePayment(payment entity.Payment, orderUID string) error
	DeletePayment(paymentID int) error
	GetPaymentsByOrderUID(orderUID string) ([]entity.Payment, error)
}

type Deliveries interface {
	CreateDelivery(delivery entity.Delivery, orderUID string) error
	GetDeliveryByOrderUID(orderUID string) ([]entity.Delivery, error)
	DeleteDelivery(deliveryID int) error
}

type Order interface {
	CreateOrder(order entity.Order) error
	DeleteOrder(orderID string) error
	GetOrderByOrderUID(orderUID string) ([]entity.Order, error)
}

type Repository struct {
	Order
	Deliveries
	Payments
	Items
}

func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{
		Order:      NewOrderPostgres(db),
		Deliveries: NewDeliveryPostgres(db),
		Payments:   NewPaymentsPostgres(db),
		Items:      NewItemsPostgres(db),
	}
}
