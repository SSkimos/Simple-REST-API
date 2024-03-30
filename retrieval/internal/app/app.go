package app

import (
	"context"
	"fmt"
	"service/internal/config"
	"service/internal/transport"
)

type App interface {
	Run(context.Context) error
}

type app struct {
	router transport.Router
	cfg    config.Config
}

func NewApp(cfg config.Config) (App, error) {
	r, err := transport.NewRouter(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create new router. %w", err)
	}

	return &app{
		router: r,
		cfg:    cfg,
	}, nil
}

func (a *app) Run(ctx context.Context) error {
	if err := a.router.Run(ctx); err != nil {
		return fmt.Errorf("failed in run. %w", err)
	}

	return nil
}
