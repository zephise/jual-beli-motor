package main

import (
	"log"

	"jual-beli-motor/repository"
	"jual-beli-motor/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repository.InitDB()

	routes.Routes()
}
