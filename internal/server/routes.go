package server

import (
	"context"
	"encoding/json"

	"github.com/ary82/balance/internal/classification"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/health", s.healthHandler)
	s.App.Get("/sse", s.sse)
	s.App.Post("/submission", s.postHandler)
	s.App.Get("/submission", s.getRandomHandler)
	s.App.Post("/t", s.temp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

func (s *FiberServer) temp(c *fiber.Ctx) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	client, err := grpc.NewClient("localhost:8000", opts...)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	serviceClient := classification.NewClassifyServiceClient(client)

	req := classification.ClassifyRequest{
		Query: []string{
			"It's an amazing day",
      "Hope you have the worst day",
      "You're special to everyone",
      "Your programming skills are weak",
		},
	}
	res, err := serviceClient.Classify(context.Background(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "in rpc call",
			"error":   err.Error(),
		})
	}

	arr := []int{}
	err = json.Unmarshal([]byte(res.Result), &arr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "in Unmarshal",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"res": arr,
	})
}
