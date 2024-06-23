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
			for {
				posMsg := html.EscapeString(s.CurrentPositivePosts.Body)
				posAuthor := html.EscapeString(s.CurrentPositivePosts.Author)
				negMsg := html.EscapeString(s.CurrentNegativePosts.Body)
				negAuthor := html.EscapeString(s.CurrentNegativePosts.Author)

				fmt.Fprintf(w,
					"event: p_b\ndata: %s\n\nevent: p_a\ndata: %s\n\nevent: n_b\ndata: %s\n\nevent: n_a\ndata: %s\n\nevent: p_c\ndata: %d\n\nevent: n_c\ndata: %d\n\n",
					posMsg, posAuthor,
					negMsg, negAuthor,
					s.PostsCount.Positive, s.PostsCount.Negative,
				)

				err := w.Flush()
				if err != nil {
					fmt.Printf("%v. Closing http connection.\n", err)
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
		return ctx.Status(fiber.StatusOK).Render(
			"form_error", fiber.Map{
				"error": err.Error(),
			},
		)
	}
	return ctx.Status(fiber.StatusOK).Render("form", nil)
}

func (s *FiberServer) index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).Render("index", nil)
}
