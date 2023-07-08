package handlers

import (
	"api/config"
	"api/database"
	"api/types"
	"context"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

// Types
type UserInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func Base(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"message": "Welcome to Shrinkr",
	})
}

func Login(c *fiber.Ctx) error {
	state := uuid.New().String()
	authConfig := config.AuthConf()
	url := authConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"url":    url,
		"state":  state,
	})
}

func GetJWT(c *fiber.Ctx) error {
	code := c.Query("code")
	authConf := config.AuthConf()
	ctx := context.Background()
	token, err := authConf.Exchange(ctx, code)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot exchange code for token",
		})
	}
	client := authConf.Client(ctx, token)
	userData, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Cannot get user data")
	}
	defer userData.Body.Close()
	body, err := ioutil.ReadAll(userData.Body)
	var userInfo UserInfo
	err = json.Unmarshal(body, &userInfo)

	var user types.User
	user.Username = userInfo.Email
	user.Joined = time.Now().Format("2006-01-02 15:04:05")

	_, err = database.GetUser(userInfo.Email)
	if err != nil {
		database.RegisterUser(&user)
	}
	claims := jwt.MapClaims{
		"email": userInfo.Email,
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Cannot create JWT token")
	}
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"token":  tokenString,
		"name":   userInfo.Name,
	})
}
