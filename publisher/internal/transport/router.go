package transport

import (
	"context"
	"fmt"
	"net"
	"service/internal/config"
	"service/internal/logger"
	"service/internal/service"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Run(context.Context) error
}

type router struct {
	eng     *gin.Engine
	service service.Service
	config  config.Router
}

func NewRouter(cfg config.Config) (Router, error) {
	s, err := service.NewService(cfg.Service)
	if err != nil {
		return nil, fmt.Errorf("failed to create new service. %w", err)
	}

	eng := initGin(cfg.Router.Debug)

	r := router{
		eng:     eng,
		service: s,
		config:  cfg.Router,
	}

	r.setupRoutes()

	return &r, nil
}

func initGin(debug bool) *gin.Engine {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	eng := gin.New()

	return eng
}

func (r *router) setupRoutes() {
	r.eng.Use(gin.Logger())

	r.eng.GET("/echo", r.echo)
}

func (r *router) Run(ctx context.Context) error {
	addr := net.JoinHostPort(r.config.IP, r.config.Port)

	logger.GetLogger().Info().Str("address", addr).Msgf("Starting endpoint on " + addr)

	if err := r.eng.Run(addr); err != nil {
		return fmt.Errorf("failed to run engine. %w", err)
	}

	return nil
}
