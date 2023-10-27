package WB_Intern_L0

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	httpServer *fiber.App
}

func (s *Server) Run(addr string, handler *fiber.App) error {
	s.httpServer = handler
	return s.httpServer.Listen(addr)
}
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.ShutdownWithContext(ctx)
}
