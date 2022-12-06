package config

import (
	"log"
	"party/database/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migration() *gorm.DB{

	dsn := "root@tcp(127.0.0.1:3306)/party2?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Migration", err)
	}

	db.Debug().AutoMigrate(
		models.Profile{},
	)
	return db
}