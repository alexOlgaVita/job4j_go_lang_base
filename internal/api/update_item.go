package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UpdateItemRequest struct {
	Name string `json:"name"`
}

func (s *Server) UpdateItem(c *fiber.Ctx) error {
	var req UpdateItemRequest
	name := c.Params("name")
	if name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}

	err := s.Repository.Update(c.Context(), name, req.Name)

	if err != nil {
		log.Errorw("s.Repository.Update", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
