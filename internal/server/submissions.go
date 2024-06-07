package server

import (
	"github.com/ary82/balance/internal/database"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) postHandler(c *fiber.Ctx) error {
	post := new(database.SubmissionReq)

	err := c.BodyParser(post)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	err = s.db.PostSubmission(post)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func (s *FiberServer) getRandomHandler(c *fiber.Ctx) error {
	submissionType := c.QueryInt("type", 0)
	if submissionType != 0 {
		submissionType = 1
	}

	post, err := s.db.GetRandom(submissionType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(post)
}
