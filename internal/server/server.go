package server

import (
	"github.com/ary82/balance/internal/post"
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	CurrentPositivePosts post.Post
	CurrentNegativePosts post.Post
	PostsCount           post.PostCounts
	App                  *fiber.App
	postService          post.PostService
}

func CreateFiberServer(
	app *fiber.App,
	postService post.PostService,
) *FiberServer {
	s := &FiberServer{
		CurrentPositivePosts: post.Post{},
		CurrentNegativePosts: post.Post{},
		PostsCount:           post.PostCounts{},
		App:                  app,
		postService:          postService,
	}

	s.RegisterRoutes()

	return s
}
