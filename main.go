package main

import (
	"api/config"
	"api/database"
	"api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// app init
	app := fiber.New()

	// connect to database
	database.ConnectRedis()
	database.ConnectMongo()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	routes.SetupRoutes(app)

	// Static files
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})
	app.Static("/404", "./public/404.html")

	app.Listen(config.Config("PORT"))
}
