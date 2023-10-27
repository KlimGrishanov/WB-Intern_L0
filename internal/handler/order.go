package handler

import (
	"WB_Intern_L0/entity"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *Handler) GetOrderById(c *fiber.Ctx) error {
	id := c.Params("id")

	delivery, err := h.services.GetDeliveryByOrderUID(id)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	payment, err := h.services.GetPaymentByOrderUID(id)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	itemsArr, err := h.services.GetItemsByOrderUID(id)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	order, err := h.services.GetOrderByOrderUID(id)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	order.Items = itemsArr
	order.Payment = payment
	order.Delivery = delivery
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
