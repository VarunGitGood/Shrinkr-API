package routes

import (
	"api/handlers"
	"api/middleware"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all routes for the application
func SetupRoutes(app *fiber.App) {
	api := app.Group("/shrinkr")
	api.Get("/", handlers.Base)
	api.Get("/login", handlers.Login)
	api.Get("/token", handlers.GetJWT)
	api.Get("/shnk/:shortURL", handlers.RedirectToLongLink)

	linksAPI := api.Group("/links")
	linksAPI.Use(middleware.AuthGuard)
	linksAPI.Post("/addURL", handlers.AddMapping)
	linksAPI.Get("/mappings", handlers.GetAllShortLinks)
	linksAPI.Delete("/:shortURL", handlers.DeleteLink)
	linksAPI.Get("/:shortURL", handlers.GetLinkById)

	userAPI := api.Group("/user")
	userAPI.Use(middleware.AuthGuard)
	api.Get("/:username", handlers.GetUser)
}
