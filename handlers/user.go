package handlers

import (
	"api/database"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Verify(c *fiber.Ctx) error {
	username := c.Request().Header.Peek("Email")
	_, err := database.GetUser(string(username))
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func GetUser(c *fiber.Ctx) error {
	username := c.Request().Header.Peek("Email")
	user, err := database.GetUser(string(username))
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}