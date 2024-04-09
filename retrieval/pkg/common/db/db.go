package db

import (
	"log"
	"retrieval/pkg/common/models/employee"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Panic("Init db error: %v", err)
	}

	db.AutoMigrate(&employee.Employee{})

	return db
}
