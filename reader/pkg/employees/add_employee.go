package employees

import (
	"gorm.io/gorm"
	"log"
	"reader/pkg/common/models"
)

func AddEmployee(db *gorm.DB, employee models.Employee) {
	if result := db.Create(&employee); result.Error != nil {
		log.Panic("Panic: %w", result.Error)
	}
}
