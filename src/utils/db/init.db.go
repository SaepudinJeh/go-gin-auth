package db

import (
	"log"
	"os"

	"github.com/SaepudinJeh/go-gin-auth/src/models/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	dbUrl := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&user.User{})
}

func GetDB() *gorm.DB {
	return db
}
