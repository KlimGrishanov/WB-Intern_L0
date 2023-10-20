package handler

import (
	"WB_Intern_L0/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *usecase.UseCase
}

func NewHandler(services *usecase.UseCase) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoute() *fiber.App {
	router := fiber.New()
	//api := router.Group("/bill-api")
	//api.Post("/income", )

	return router
}
