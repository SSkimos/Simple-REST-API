package employees

import (
	"gorm.io/gorm"
	"reader/pkg/common/models"
)

func CheckExistingID(db *gorm.DB, employee models.Employee) bool {
	existingData := models.Employee{}

	if err := db.Where("id = ?", employee.Id).First(&existingData).Error; err == nil {
		// ID уже существует в базе данных
		return true
	}
	return false
}
