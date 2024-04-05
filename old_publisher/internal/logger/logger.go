package logger

import (
	"fmt"
	"os"
	"service/internal/config"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

const diodeMsgsCount = 1000

var loggertInstance *zerolog.Logger

func GetLogger() *zerolog.Logger {
	return loggertInstance
}

func InitLogger(serviceTag string, cfg config.Logger) {
	w := diode.NewWriter(os.Stderr, diodeMsgsCount, 0, func(missed int) {
		fmt.Printf("Dropped %d messages", missed)
	})
	rt := zerolog.New(w)

	isTerm := isatty.IsTerminal(os.Stderr.Fd())

	if isTerm {
		w := zerolog.ConsoleWriter{Out: os.Stderr}
		w.NoColor = !isTerm
		rt = zerolog.New(w)
	}

	lvl := zerolog.InfoLevel

	if cfg.Level != "" {
		var err error
		lvl, err = zerolog.ParseLevel(cfg.Level)

		if err != nil {
			GetLogger().Err(err).Msgf("Failed to parser log level: %v. Use %v", cfg.Level, lvl.String())
		}
	}

	rt = rt.Level(lvl).With().Timestamp().Logger().With().Str(zerolog.CallerFieldName, serviceTag).Logger()
	rt.Info().Str("level", lvl.String()).Msg("Setup log level")
	loggertInstance = &rt
}
