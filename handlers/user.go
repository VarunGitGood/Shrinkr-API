package handlers

import (
	"api/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	username := c.Params("username")
	user, err := database.GetUser(username)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot get user",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}
