package handler

import "github.com/gofiber/fiber/v2"

func (h *Handler) Example(c *fiber.Ctx) error {
	return c.SendString("I'm a GET request!")
}
