package handlers

import (
	"api/config"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

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

// working
func Base(c *fiber.Ctx) error {
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status":  "success",
		"message": "Welcome to Shrinkr",
	})
}

// OAuth2
func Login(c *fiber.Ctx) error {
	// send a string containing the URL to redirect to for login
	// set the state to a unique string to identiffy user later on ClI side
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
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Cannot unmarshal user data")
	}
	// TODO login user if exists, else register user
	claims := jwt.MapClaims{
		"email": userInfo.Email,
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Config("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Cannot create JWT token")
	}
	fmt.Println(tokenString)
	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"token":  tokenString,
	})
}
