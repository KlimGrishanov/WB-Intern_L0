package repo

import (
	"WB_Intern_L0/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ItemsPostgres struct {
	db *sqlx.DB
}

func (i ItemsPostgres) CreateItem(item entity.Item, orderUID string) error {
	query := fmt.Sprintf("INSERT INTO %s (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)"+" values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", itemsTable)
	_, err := i.db.Exec(query, orderUID, item.ChrtId, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status)
	return err
}

func (i ItemsPostgres) GetItemsByOrderUID(orderUID string) ([]entity.Item, error) {
	var items []entity.Item

	query := fmt.Sprintf("SELECT * FROM %s WHERE order_uid=$1", itemsTable)
	err := i.db.Select(&items, query, orderUID)

	return items, err
}

func (i ItemsPostgres) DeleteItem(itemID int) error {
	//TODO implement me
	panic("implement me")
}

func NewItemsPostgres(db *sqlx.DB) *ItemsPostgres {
	return &ItemsPostgres{db: db}
}
