package models

type Employee struct {
	EmployeeId  int    `gorm:"unique"`
	FirstName   string `json:"first_name"`
	Last_name   string `json:"last_name"`
	Position    string `json:"position"`
	Department  string `json:"department"`
	HireDate    string `json:"hire_date"`
	Salary      int    `json:"salary"`
	Email       string `json:"email"`
	PhoneNumber int    `json:"phone_number"`
	Address     string `json:"address"`
}
