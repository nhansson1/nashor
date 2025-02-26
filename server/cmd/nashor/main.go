package main

import (
	"log"
	"nashor/internal/server"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()

    if err != nil {
        log.Fatal("Failed to load .env file!")
    }

    server.Run()
}
