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
	api.Post("/create/", h.CreateOrder)
}

// CreateOrder - создание заказа
// @Summary Create a new order
// @Description This endpoint allows you to create a new order
// @Tags order
// @Accept json
// @Produce json
// @Param order body models.Order true "Order creation data"
// @Success 201 {object} fiber.Map{"message": string, "order_id": uint} "Order successfully created"
// @Failure 400 {object} fiber.Map{"error": string} "Invalid request format"
// @Failure 500 {object} fiber.Map{"error": string} "Failed to create order"
// @Router /order/register [post]
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
