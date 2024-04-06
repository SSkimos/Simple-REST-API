package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDBConnection(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal("Init db error: %v", err)
	}

	return db
}
