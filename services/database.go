package services

import (
	"go-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open(sqlite.Open("go-api.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Book{})

	DB = database
}
