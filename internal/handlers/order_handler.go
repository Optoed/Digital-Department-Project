package handlers

import (
	"backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &CourierHandler{service: service}
}

func OrderRoutes(api fiber.Router) {
	api.Post("/order/", OrderRegister)
}

// TODO
func OrderRegister(c *fiber.Ctx) error {
	return nil
}

// TODO
func OrderGetCoordinates(c *fiber.Ctx) {

}

// TODO
func OrderPayCost(c fiber.Ctx) {

}
