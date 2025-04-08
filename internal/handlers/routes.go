package handlers

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(api fiber.Router) {
	api.Get("/", helloHandler)
}

func helloHandler(c *fiber.Ctx) error {
	return c.SendString("Привет из Fiber 👋")
}
