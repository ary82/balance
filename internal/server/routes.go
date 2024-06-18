package server

import (
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (s *FiberServer) RegisterRoutes() {
	s.App.Use(logger.New())
	s.App.Use(compress.New())
	s.App.Use(favicon.New())

	s.App.Get("/sse", s.sse)
	s.App.Get("/", s.index)
	s.App.Post("/", s.submitPost)
}
