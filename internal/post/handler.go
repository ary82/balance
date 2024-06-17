package post

import (
	"bufio"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type PostHandler struct {
	postService PostService
	sseCh       chan string
}

func NewPostHandler(postRouter fiber.Router, messageCh chan string, postService PostService) {
	handler := &PostHandler{
		postService: postService,
		sseCh:       messageCh,
	}

	postRouter.Get("/test", handler.test)
	postRouter.Get("/sse", handler.sse)
	postRouter.Get("/:type", handler.getRandomPost)
	postRouter.Post("/", handler.postPost)
}

func (h *PostHandler) test(ctx *fiber.Ctx) error {
	return ctx.Render("index", 0)
}

func (h *PostHandler) getRandomPost(ctx *fiber.Ctx) error {
	postType, err := ctx.ParamsInt("type")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	post, err := h.postService.GetRandomPost(postType)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"post":    post,
	})
}

func (h *PostHandler) postPost(ctx *fiber.Ctx) error {
	post := new(Post)

	err := ctx.BodyParser(post)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err = h.postService.CreatePost(post)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func (h *PostHandler) sse(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(
		fasthttp.StreamWriter(func(w *bufio.Writer) {
			fmt.Println("WRITER")

			for msg := range h.sseCh {
				fmt.Fprintf(w, "data: Message: %s\n\n", msg)
				fmt.Println(msg)

				err := w.Flush()
				if err != nil {
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
					break
				}
				time.Sleep(2 * time.Second)
			}
		}),
	)

	return nil
}
