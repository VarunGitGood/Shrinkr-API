package middleware

import (
	"api/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// AuthGuard is a middleware that checks if the user is authenticated
func AuthGuard(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}
	tokenString = tokenString[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}
	claims := token.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	c.Request().Header.Set("UserEmail", email)
	return c.Next()
}
