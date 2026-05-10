package api

import "github.com/gofiber/fiber/v2"

func (s *Server) Route(route fiber.Router) {
	route.Post("/item/", s.CreateItem)
	route.Get("/items/", s.GetItems)
	route.Get("/item/:name", s.GetItem)
	route.Put("/item/:name", s.UpdateItem)
	route.Delete("/item/:name", s.UpdateItem)
}
