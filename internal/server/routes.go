package server

import (
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (s *FiberServer) RegisterRoutes() {
	s.App.Use(logger.New())
	s.App.Use(compress.New())
	s.App.Use(favicon.New(favicon.Config{
		File: "./static/icons/favicon.png",
	}))

	s.App.Static("/main.css", "./static/styles/output.css")
	s.App.Static("/htmx.min.js", "./static/js/htmx.min.js")
	s.App.Static("/sse.min.js", "./static/js/sse.min.js")

	s.App.Get("/sse", s.sse)
	s.App.Get("/", s.index)
	s.App.Post("/", s.submitPost)
}
