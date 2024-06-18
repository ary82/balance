package server

import (
	"bufio"
	"fmt"
	"html"
	"time"

	"github.com/ary82/balance/internal/post"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func (s *FiberServer) sse(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")

	c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(
		fasthttp.StreamWriter(func(w *bufio.Writer) {
			fmt.Println("WRITER")

			for {
				posMsg := html.EscapeString(s.CurrentPositivePosts.Body)
				posAuthor := html.EscapeString(s.CurrentPositivePosts.Author)
				negMsg := html.EscapeString(s.CurrentNegativePosts.Body)
				negAuthor := html.EscapeString(s.CurrentNegativePosts.Author)

				_ = posAuthor
				_ = negAuthor

				fmt.Fprintf(w,
					"event: positive_body\ndata: %s\n\nevent: positive_author\ndata: %s\n\nevent: negative_body\ndata: %s\n\nevent: negative_author\ndata: %s\n\n",
					posMsg, posAuthor,
					negMsg, negAuthor,
				)

				err := w.Flush()
				if err != nil {
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
					break
				}
				time.Sleep(3 * time.Second)
			}
		}),
	)

	return nil
}

func (s *FiberServer) submitPost(ctx *fiber.Ctx) error {
	body := ctx.FormValue("body")
	author := ctx.FormValue("author")

	post := &post.Post{
		Body:   body,
		Author: author,
	}

	err := s.postService.CreatePost(post)
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

func (s *FiberServer) index(ctx *fiber.Ctx) error {
	return ctx.Render("index", nil)
}
