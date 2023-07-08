package handlers

import (
	"api/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AddMapping(c *fiber.Ctx) error {
	username := c.Request().Header.Peek("Email")
	link := new(database.LinkDTO)
	c.BodyParser(link)
	if link.ShortURL == "" || link.LongURL == "" || link.Description == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Missing required fields",
		})
	}
	err := database.AddURL(link, string(username))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot add mapping",
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

func GetAllShortLinks(c *fiber.Ctx) error {
	username := c.Request().Header.Peek("Email")
	mappings, err := database.GetUrlsByUser(string(username))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot get links",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   mappings,
	})
}

func DeleteLink(c *fiber.Ctx) error {
	username := c.Request().Header.Peek("Email")
	shortURL := c.Params("shortURL")
	err := database.DeleteLink(shortURL, string(username))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot delete link",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Link deleted",
	})
}

func GetLinkById(c *fiber.Ctx) error {
	username := c.Request().Header.Peek("Email")
	shortURL := c.Params("shortURL")
	mapping, err := database.GetLinkInfo(shortURL, string(username))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot get link",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   mapping,
	})
}

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
