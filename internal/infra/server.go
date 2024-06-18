package infra

import (
	"database/sql"

	"github.com/ary82/balance/internal/post"
	"github.com/ary82/balance/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func NewFiberServer(db *sql.DB) *server.FiberServer {
	// Initialize standard Go html template engine
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		AppName:      "Balance",
		ServerHeader: "Fiber",
		Views:        engine,
	})

	postRepo := post.NewPostRepository(db)
	postservice := post.CreatePostService(postRepo)

	server := server.CreateFiberServer(app, postservice)
	return server
}
