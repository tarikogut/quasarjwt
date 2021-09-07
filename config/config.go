package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	// load .env file
	err := godotenv.Load("app.conf")
	if err != nil {
		fmt.Print("Error loading config file file")
	}
	return os.Getenv(key)
}