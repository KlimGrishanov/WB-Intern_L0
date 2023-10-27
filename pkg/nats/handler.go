package natsService

import (
	"WB_Intern_L0/entity"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func (nc *Nats) GetOrderByIdNATS(m *nats.Msg) {
	var id string
	id = string(m.Data)

	order, isExist := nc.Cache.GetOrder(id)
	if isExist == false {
		nc.NATSServer.Publish("OrderInfo", []byte("No order"))
	} else {
		orderMsg, err := json.Marshal(order)

		if err != nil {
			zap.L().Error(err.Error())
			nc.NATSServer.Publish("OrderInfo", []byte("Error"))
		} else {
			nc.NATSServer.Publish("OrderInfo", orderMsg)
		}
	}

}

func (nc *Nats) CreateOrderNATS(m *nats.Msg) {
	var order entity.Order

	err := json.Unmarshal(m.Data, &order)
	if err != nil {
		zap.L().Error("can't unmarshal data into struct")
		return
	}

	nc.Cache.NewOrder(order)

	err = nc.Services.NewOrder(order)
	if err != nil {
		zap.L().Error(err.Error())
	}

	err = nc.Services.NewPayment(order.Payment, order.OrderUid)
	if err != nil {
		zap.L().Error(err.Error())
	}

	err = nc.Services.NewDelivery(order.Delivery, order.OrderUid)
	if err != nil {
		zap.L().Error(err.Error())
	}

	for i := 0; i < len(order.Items); i++ {
		err = nc.Services.NewItem(order.Items[i], order.OrderUid)
		if err != nil {
			zap.L().Error(err.Error())
		}
	}
}
