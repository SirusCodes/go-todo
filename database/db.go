package database

import (
	"fmt"

	"github.com/SirusCodes/go-todo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to DB")
		panic(err.Error())
	}

	DBConn = db

	sqlDB, _ := db.DB()

	sqlDB.Stats()

	db.AutoMigrate(&models.Todo{})
	db.AutoMigrate(&models.User{})
}
