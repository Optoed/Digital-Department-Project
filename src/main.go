package main

import (
	"backend/internal/db"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"backend/internal/handlers"
)

func main() {
	// Загружаем переменные окружения
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Подключение к БД (например, PostgreSQL)
	if err := db.Init(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Создаем Fiber-приложение
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	// Роуты
	api := app.Group("/api")
	handlers.RegisterRoutes(api)

	// Запуск
	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(app.Listen(":" + port))
}
