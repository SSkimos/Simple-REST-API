package employees

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"publisher/pkg/common/models"
	"time"
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
	nc, err := nats.Connect("stan-nats:4222")
	if err != nil {
		log.Panic(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer nc.Close()

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
	employee.Address = body.Address

	jsonData, err := json.Marshal(employee)
	if err != nil {
		log.Panic("Json marshall error: %w", err)
	}
	subj, msg := "new_employee", jsonData

	timeout := 60 * time.Second
	response, err := nc.Request("new_employee", jsonData, timeout)

	if err != nil {
		if err == nats.ErrTimeout {
			log.Printf("Error: request timed out after %v", timeout)
		} else {
			log.Printf("Panic: %w", err)
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	nc.Flush()

	log.Printf("Published [%s] : '%s'\n", subj, msg)
	log.Printf("Received response [%s] : '%s'\n", string(response.Data))
	c.JSON(http.StatusOK, &employee)
}
