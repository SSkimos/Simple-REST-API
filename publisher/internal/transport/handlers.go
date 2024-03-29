package transport

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
