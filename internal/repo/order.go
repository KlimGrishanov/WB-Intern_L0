package repo

import (
	"WB_Intern_L0/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func (o OrderPostgres) CreateOrder(order entity.Order) error {
	query := fmt.Sprintf("INSERT INTO %s (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)"+" values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", ordersTable)
	_, err := o.db.Exec(query, order.OrderUid, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	return nil
}

func (o OrderPostgres) DeleteOrder(orderID string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE order_uid", ordersTable)
	_, err := o.db.Exec(query, orderID)
	return err
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}
