package employees

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"retrieval/pkg/common/models"
)

func (h handler) GetEmployees(c *gin.Context) {
	var employees []models.Employee

	if result := h.DB.Find(&employees); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &employees)
}
