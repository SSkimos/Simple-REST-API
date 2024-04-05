package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"service/internal/models"
)

func (r *router) echo(c *gin.Context) {
	number := c.DefaultQuery("number", "0")
	v, err := strconv.Atoi(number)
	if err != nil {
		c.JSON(http.StatusBadRequest, "where is my mind¿")

		return
	}

	response := r.service.Echo(v)

	c.JSON(http.StatusOK, response)
}

func (r *router) employeeAddHandler(c *gin.Context) {
	var request models.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	response := r.service.EmployeeAdd(&request)

	c.JSON(http.StatusOK, response)
}
