package routes

import (
	"github.com/gofiber/fiber/v2"
	"api/handlers"
)

// SetupRoutes sets up all routes for the application
func SetupRoutes(app *fiber.App) {
	api := app.Group("/shrinkr")
	// routes
	api.Get("/", handlers.Base)
	api.Get("/login", handlers.Login)
	api.Get("/token", handlers.GetJWT)
	api.Post("/addURL/:username", handlers.AddMapping)
	api.Get("/mappings/:username", handlers.GetAllShortLinks)
}
