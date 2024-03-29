package commands

import (
	"context"
	"fmt"
	"service/internal/app"
	"service/internal/common"
	"service/internal/config"
	"service/internal/logger"
)

func Execute() {
	cfg, err := FindConfig[config.Config]()
	if err != nil {
		fmt.Println("failed to find config. " + err.Error())

		return
	}

	logger.InitLogger(common.ServiceName, cfg.Logger)

	printConfig[config.Config](cfg)

	a, err := app.NewApp(cfg)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to create new app")

		return
	}

	if err := a.Run(logger.GetLogger().WithContext(context.Background())); err != nil {
		logger.GetLogger().Err(err).Msg("failed into app run")

		return
	}
}
