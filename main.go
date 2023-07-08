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
	app.Use(recover.New())

	// Routes
	routes.SetupRoutes(app)

	// Static files
	// add one for if link is password protected
	// add one for if link is not found

	app.Listen(config.Config("PORT"))
}
