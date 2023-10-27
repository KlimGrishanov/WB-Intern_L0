package handler

import (
	"WB_Intern_L0/entity"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *Handler) GetOrderById(c *fiber.Ctx) error {
	return c.SendString("I'm a GET request!")
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
