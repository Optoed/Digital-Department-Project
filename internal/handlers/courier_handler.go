package handlers

import (
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type CourierHandler struct {
	service *services.CourierService
}

func NewCourierHandler(service *services.CourierService) *CourierHandler {
	return &CourierHandler{service: service}
}

func (h *CourierHandler) CourierRoutes(api fiber.Router) {
	api.Post("/register/", h.Register)
	api.Post("/login/", h.Login)

	protected := api.Group("/secure", middleware.JWTMiddleware())
	protected.Post("/take_order/", h.TakeOrder)
	protected.Post("/finish_delivery/", h.FinishDelivery)
	protected.Get("/get_directions/", h.GetShortestDirections)
	protected.Post("/rate/", h.Rate)
}

// Register - регистрация курьера
func (h *CourierHandler) Register(c *fiber.Ctx) error {
	var courier models.Courier
	if err := c.BodyParser(&courier); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request body",
		})
	}

	id, err := h.service.Register(&courier)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register courier",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

// Login - авторизация курьера
func (h *CourierHandler) Login(c *fiber.Ctx) error {
	// Пример авторизации (например, проверка email/phone и пароля)
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse login request",
		})
	}

	token, err := h.service.AuthenticateCourier(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

// TakeOrder - курьер берет заказ
func (h *CourierHandler) TakeOrder(c *fiber.Ctx) error {
	var orderRequest struct {
		CourierID uint `json:"courier_id"`
		OrderID   uint `json:"order_id"`
	}

	if err := c.BodyParser(&orderRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.service.TakeOrder(orderRequest.CourierID, orderRequest.OrderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to take order",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order taken successfully",
	})
}

// FinishDelivery - завершение доставки
func (h *CourierHandler) FinishDelivery(c *fiber.Ctx) error {
	var deliveryRequest struct {
		CourierID uint `json:"courier_id"`
		OrderID   uint `json:"order_id"`
	}

	if err := c.BodyParser(&deliveryRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.service.FinishDelivery(deliveryRequest.CourierID, deliveryRequest.OrderID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to finish delivery",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delivery finished successfully",
	})
}

// GetShortestDirections - получение кратчайшего маршрута
func (h *CourierHandler) GetShortestDirections(c *fiber.Ctx) error {
	courierIDstr := c.Query("courier_id")
	courierID, err := strconv.Atoi(courierIDstr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "courierID is not a number",
		})
	}

	orderIDstr := c.Query("order_id")
	orderID, err := strconv.Atoi(orderIDstr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "orderID is not a number",
		})
	}

	// Заглушка: получаем маршрут из сервиса
	// Если интеграция с внешним API (например, Яндекс Карты), тут можно сделать вызов
	direction, err := h.service.GetDirections(uint(courierID), uint(orderID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get directions",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"directions": direction,
	})
}

// Rate - оценка курьера
func (h *CourierHandler) Rate(c *fiber.Ctx) error {
	var ratingRequest struct {
		CourierID uint    `json:"courier_id"`
		Rating    float64 `json:"rating"`
	}

	if err := c.BodyParser(&ratingRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := h.service.RateCourier(ratingRequest.CourierID, ratingRequest.Rating)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to rate courier",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Courier rated successfully",
	})
}
