package employees

import (
	"gorm.io/gorm"
	"reader/pkg/common/models/employee"
)

func CheckExistingID(db *gorm.DB, newEmployee employee.Employee) bool {
	var employee employee.Employee

	if err := db.Where("id = ?", newEmployee.Id).First(&employee).Error; err == nil {
		// ID уже существует в базе данных
		return true
	}
	return false
}
