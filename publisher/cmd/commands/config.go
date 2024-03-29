package commands

import (
	"errors"
	"fmt"
	"io"
	"os"
	"service/internal/logger"

	"gopkg.in/yaml.v3"
)

var ErrMissingConfigFlag = errors.New("missing config flag")
var ErrMissingConfigPath = errors.New("missing config path")

func FindConfig[T any]() (T, error) {
	var (
		cfg       T
		flagIndex int
	)

	for i, v := range os.Args {
		if v == "--config" {
			flagIndex = i
		}
	}

	if flagIndex == 0 {
		return cfg, ErrMissingConfigFlag
	}

	if len(os.Args)-1 == flagIndex {
		return cfg, ErrMissingConfigPath
	}

	path := os.Args[flagIndex+1]

	f, err := os.Open(path)
	if err != nil {
		return cfg, fmt.Errorf("failed to open file. %w", err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return cfg, fmt.Errorf("failed to read file. %w", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("failed to unmarhsal. %w", err)
	}

	return cfg, nil
}

func printConfig[T any](cfg T) {
	cfgData, err := yaml.Marshal(cfg)
	if err != nil {
		logger.GetLogger().Debug().Err(err).Msg("Could not marshal configuration")

		return
	}

	logger.GetLogger().Debug().Msg(string(cfgData))
}
