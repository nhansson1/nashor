package main

import (
	"github.com/joho/godotenv"
	"log"
	"nashor/internal/server"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load .env file!")
	}

	server.Run()
}
