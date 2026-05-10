package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (s *Server) DeleteItem(c *fiber.Ctx) error {
	name := c.Params("name")
	if name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}

	err := s.Repository.Delete(c.Context(), name)

	if err != nil {
		log.Errorw("s.Repository.Delete", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
