package handlers

import "github.com/gofiber/fiber/v2"

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
