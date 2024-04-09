package employees

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"log"
	"net/http"
	"publisher/pkg/common/models/employee"
	"time"
)

func (h handler) AddEmployee(c *gin.Context) {
	var employee employee.Employee

	if err := c.BindJSON(&employee); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	jsonData, err := json.Marshal(employee)
	if err != nil {
		log.Panic("Json marshall error: %w", err)
	}
	subj, msg := "new_employee", jsonData

	timeout := 60 * time.Second
	response, err := h.nc.Request("new_employee", jsonData, timeout)
	log.Printf("Published [%s] : '%s'\n", subj, msg)

	if err != nil {
		if err == nats.ErrTimeout {
			log.Printf("Error: request timed out after %v", timeout)
		} else {
			log.Printf("Panic: %w", err)
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	h.nc.Flush()

	switch string(response.Data) {
	case "201 Created":
		log.Printf("Received response [%s] : '%s'\n", string(response.Data))
		c.JSON(http.StatusNoContent, nil)

	case "409 Conflict":
		log.Printf("Received response [%s] : '%s'\n", string(response.Data))
		c.JSON(http.StatusConflict, nil)
	}
}
