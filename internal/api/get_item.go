package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

//
//type ItemRequest struct {
//	ID   string `json:"id"`
//	Name string `json:"name"`
//}

type GetItemResponse struct {
	Item ItemRequest `json:"item"`
}

func (s *Server) GetItem(c *fiber.Ctx) error {
	name := c.Params("name")
	if name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}

	item, err := s.Repository.Get(c.Context(), name)
	if err != nil {
		log.Errorw("s.Repository.Get", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	res := ItemRequest{
		ID:   item.ID,
		Name: item.Name,
	}

	return c.Status(fiber.StatusOK).JSON(GetItemResponse{Item: res})
}
