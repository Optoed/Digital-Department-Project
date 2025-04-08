package handlers

import "github.com/gofiber/fiber/v2"

func CourierRoutes(api fiber.Router) {
	api.Post("/courier/", CourierRegister)
}

// TODO
func CourierRegister(c *fiber.Ctx) error {
	return nil
}

// TODO
func CourierAddOrder(c *fiber.Ctx) error {
	return nil
}

// TODO
func CourierRate(c *fiber.Ctx) error {
	return nil
}
