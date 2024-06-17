package infra

import (
	"database/sql"

	"github.com/ary82/balance/internal/post"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiberServer(db *sql.DB, sseCh chan string) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "Balance",
		ServerHeader: "Fiber",
	})

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(favicon.New())

	postRepo := post.NewPostRepository(db)
	postservice := post.CreatePostService(postRepo)

	post.NewPostHandler(app.Group("/post"), sseCh, postservice)

	return app
}
