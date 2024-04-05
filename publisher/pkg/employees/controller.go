package employees

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

type handler struct {
	con *nats.Conn
}

func RegisterRoutes(r *gin.Engine, c *nats.Conn) {
	h := &handler{
		con: c,
	}

	routes := r.Group("/employees")
	routes.POST("/", h.AddEmployee)
}
