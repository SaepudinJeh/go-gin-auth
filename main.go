package main

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error load .env file")
	}
}

func main() {
	println("asdsa")
}
