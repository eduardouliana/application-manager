package database

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	databaseName := "gerapp.db"

	database, err := gorm.Open(sqlite.Open(databaseName), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database! Error: " + err.Error())
	}

	db = database

	log.Print("Database connected!")

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	return db
}
