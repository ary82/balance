package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/health", s.healthHandler)
	s.App.Get("/sse", s.sse)
	s.App.Post("/submission", s.postHandler)
	s.App.Get("/submission", s.getRandomHandler)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
