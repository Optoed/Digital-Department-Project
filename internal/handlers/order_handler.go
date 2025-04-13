package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

func (h *OrderHandler) OrderRoutes(api fiber.Router) {
	api.Post("/register/", h.CreateOrder)
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var input models.Order

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "неверный формат запроса",
		})
	}

	orderID, err := h.service.Create(&input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "не удалось создать заказ",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "заказ успешно создан",
		"order_id": orderID,
	})
}
