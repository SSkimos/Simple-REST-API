package employees

import (
	"github.com/curtrika/reader/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddEmployeeRequestBody struct {
	EmployeeId  int    `json:"employee_id"`
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

func (h handler) AddEmployee(c *gin.Context) {
	body := AddEmployeeRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var employee models.Employee

	employee.EmployeeId = body.EmployeeId
	employee.FirstName = body.FirstName
	employee.Last_name = body.Last_name
	employee.Position = body.Position
	employee.Department = body.Department
	employee.HireDate = body.HireDate
	employee.Salary = body.Salary
	employee.Email = body.Email
	employee.PhoneNumber = body.PhoneNumber
	employee.Address = employee.Address

	if result := h.DB.Create(&employee); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &employee)
}
