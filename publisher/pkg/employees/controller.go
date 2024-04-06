package employees

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

type handler struct {
	nc *nats.Conn
}

func RegisterRoutes(r *gin.Engine, ncon *nats.Conn) {
	h := &handler{
		nc: ncon,
	}

	routes := r.Group("/employees")
	routes.POST("/", h.AddEmployee)
}
