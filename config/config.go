package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"golang.org/x/oauth2"
)

func AuthConf() oauth2.Config {
	authConfig := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile","https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  os.Getenv("AUTH_URL"),
			TokenURL: os.Getenv("TOKEN_URL"),
		},
		RedirectURL: os.Getenv("REDIRECT_URI"),
	}
	return *authConfig
}

func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
