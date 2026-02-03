package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//go:embed README.md
var readmeContent string

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "SecondBrainManager"
	}

	fmt.Printf("Welcome to %s!\n", appName)
}
