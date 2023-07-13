package handlers

import (
	"api/database"
	"api/types"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AddMapping(c *fiber.Ctx) error {
	username := c.Request().Header.Peek("Email")
	link := new(types.LinkDTO)
	c.BodyParser(link)
	error := link.Validate()
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": error.Error(),
		})
	}
	err := database.AddURL(link, string(username))
	if err != nil {
		err := err.(*types.CustomError)
		return c.Status(err.StatusCode()).JSON(fiber.Map{
			"status":  err.StatusCode(),
			"message": err.Error(),
		})
	}
	if err := database.StoreMapping(link); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot store mapping",
		})
	}
	fmt.Println("Mapping stored")
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
			"message": err.Error(),
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
	longURL, err := database.GetLongURL(shortURL)
	if err != nil {
		fmt.Printf(longURL)
		fmt.Println(err.Error())
		return c.Redirect("/404")
	}
	return c.Redirect(longURL)
}
