package handler

import (
	"WB_Intern_L0/entity"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *Handler) GetOrderById(c *fiber.Ctx) error {
	id := c.Params("id")

	order, isExist := h.cache.GetOrder(id)
	if isExist == false {
		zap.L().Error("order doesn't exist")
		return c.Send([]byte("order doesn't exist"))
	}

	deliveryJSON, err := json.Marshal(order)

	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	return c.Send(deliveryJSON)
}

func (h *Handler) CreateOrder(c *fiber.Ctx) error {
	order := new(entity.Order)

	if err := c.BodyParser(order); err != nil {
		zap.L().Error(err.Error())
		return err
	}

	h.cache.NewOrder(*order)

	err := h.services.NewOrder(*order)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	err = h.services.NewPayment(order.Payment, order.OrderUid)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	err = h.services.NewDelivery(order.Delivery, order.OrderUid)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	for i := 0; i < len(order.Items); i++ {
		err = h.services.NewItem(order.Items[i], order.OrderUid)
		if err != nil {
			zap.L().Error(err.Error())
			return err
		}
	}

	c.SendStatus(200)
	return nil
}
