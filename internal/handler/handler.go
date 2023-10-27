package handler

import (
	"WB_Intern_L0/internal/usecase"
	"WB_Intern_L0/pkg/cache"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	cache    *cache.CacheService
	services *usecase.UseCase
}

func NewHandler(services *usecase.UseCase, cache *cache.CacheService) *Handler {
	return &Handler{services: services, cache: cache}
}

func (h *Handler) InitRoute() *fiber.App {
	router := fiber.New()
	api := router.Group("/api")
	api.Get("/order/:id", h.GetOrderById)
	api.Post("/order", h.CreateOrder)
	return router
}
