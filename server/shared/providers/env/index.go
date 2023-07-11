package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type envVariables string

const (
	Port    envVariables = "PORT"
	BaseURL envVariables = "BASE_URL"
	DbURL   envVariables = "DB_URL"
)

func Load() {
	var err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file. Please check if you have one")
		os.Exit(1)
	}
}

func Get(s envVariables) string {
	var env = os.Getenv(string(s))
	if env == "" {
		fmt.Printf("Variable %v not found. Make sure you have it in the .env file.\n", string(s))
		os.Exit(1)
	}

	return env
}
