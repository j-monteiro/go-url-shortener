package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvConfig()
}

func loadEnvConfig() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
