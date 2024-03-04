package main

import (
	"log"

	"github.com/SaepudinJeh/go-gin-auth/src/utils/db"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error load .env file")
	}

	db.Init()
}

func main() {
	println("asdsa")
}
