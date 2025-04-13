package main

import (
	"backend/internal/db"
	"backend/internal/handlers"
	"backend/internal/repository"
	"backend/internal/services"
	"backend/internal/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Загружаем переменные окружения
	port := os.Getenv("PORT")
	utils.SetJwtSecret()

	// Подключение к БД (например, PostgreSQL)
	if err := db.Init(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Создаем Fiber-приложение
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// repo, service, handler
	orderRepo := repository.NewOrderRepo(db.DB)
	courierRepo := repository.NewCourierRepo(db.DB)

	orderService := services.NewOrderService(orderRepo)
	courierService := services.NewCourierService(courierRepo)

	orderHandler := handlers.NewOrderHandler(orderService)
	courierHandler := handlers.NewCourierHandler(courierService)

	// Роуты
	api := app.Group("/api")

	courierApi := api.Group("/courier")
	courierHandler.CourierRoutes(courierApi)

	orderApi := api.Group("/order")
	orderHandler.OrderRoutes(orderApi)

	// Запуск
	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(app.Listen(":" + port))
}
