package db

import (
	"log"
	"retrieval/pkg/common/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Panic("Init db error: %v", err)
	}

	db.AutoMigrate(&models.Employee{})

	return db
}
