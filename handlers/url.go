package handlers

import (
	"api/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// store the mapping of short link to long link (short link is generated from CLI)
func AddMapping(c *fiber.Ctx) error {
	// username := c.Request().Header.Peek("Email")
	link := new(database.Link)
	c.BodyParser(link)
	if link.ShortURL == "" || link.LongURL == "" || link.Description == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing required fields",
		})
	}
	if err := database.StoreMapping(link); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot store mapping",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Mapping stored",
	})
}

// get list of all short links for a user
func GetAllShortLinks(c *fiber.Ctx) error {
	username := c.Request().Header.Peek("Email")
	mappings, err := database.GetMappings(string(username))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot get links",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   mappings,
	})
}

// redirect to long link
func RedirectToLongLink(c *fiber.Ctx) error {
	shortURL := c.Params("shortURL")
	username := c.Request().Header.Peek("Email")
	fmt.Println(string(username))
	longURL, err := database.GetLongURL(shortURL)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot get long URL",
		})
	}
	return c.Redirect(longURL)
}
