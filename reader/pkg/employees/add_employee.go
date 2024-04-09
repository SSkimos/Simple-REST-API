package employees

import (
	"gorm.io/gorm"
	"log"
	"reader/pkg/common/models/employee"
)

func AddEmployee(db *gorm.DB, employee employee.Employee) {
	if result := db.Create(&employee); result.Error != nil {
		log.Panic("Panic: %w", result.Error)
	}
}
