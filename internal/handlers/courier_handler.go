package handlers

import (
	"backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

type CourierHandler struct {
	service *services.CourierService
}

func NewCourierHandler(service *services.CourierService) *CourierHandler {
	return &CourierHandler{service: service}
}

func (h *CourierHandler) CourierRoutes(api fiber.Router) {
	api.Post("/courier/register/", h.CourierRegister)
}

// TODO
func (h *CourierHandler) CourierRegister(c *fiber.Ctx) error {

}

// TODO
func CourierTakeOrder(c *fiber.Ctx) error {

}

// TODO
func CourierFinishDelivery(c *fiber.Ctx) {

}

// TODO
func Courier() {

}

// TODO
func CourierRate(c *fiber.Ctx) error {

}
