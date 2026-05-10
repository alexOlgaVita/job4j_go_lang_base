package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UpdateItemRequest struct {
	Name string `json:"name"`
}

type UpdateItemResponse struct {
	Item ItemRequest `json:"item"`
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

	item, err := s.Repository.Get(c.Context(), name)
	if err != nil {
		log.Errorw("s.Repository.Get", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	err = s.Repository.Update(c.Context(), name, req.Name)

	if err != nil {
		log.Errorw("s.Repository.Update", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	res := ItemRequest{
		ID:   item.ID,
		Name: req.Name,
	}
	return c.Status(fiber.StatusOK).JSON(UpdateItemResponse{Item: res})
}
