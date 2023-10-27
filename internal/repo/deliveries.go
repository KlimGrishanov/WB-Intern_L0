package repo

import (
	"WB_Intern_L0/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DeliveriesPostgres struct {
	db *sqlx.DB
}

func (d DeliveriesPostgres) CreateDelivery(delivery entity.Delivery, orderUID string) error {
	query := fmt.Sprintf("INSERT INTO %s (order_uid, name, phone, zip, city, address, region, email)"+" values($1, $2, $3, $4, $5, $6, $7, $8)", deliveriesTable)
	_, err := d.db.Exec(query, orderUID, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	return err
}

func (d DeliveriesPostgres) GetDeliveryByOrderUID(orderUID string) ([]entity.Delivery, error) {
	var delivery []entity.Delivery

	query := fmt.Sprintf("SELECT * FROM %s WHERE order_uid=$1", deliveriesTable)
	err := d.db.Select(&delivery, query, orderUID)

	return delivery, err
}

func (d DeliveriesPostgres) DeleteDelivery(orderUID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE order_uid", deliveriesTable)
	_, err := d.db.Exec(query, orderUID)
	return err
}

func NewDeliveryPostgres(db *sqlx.DB) *DeliveriesPostgres {
	return &DeliveriesPostgres{db: db}
}
