package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var LOG zerolog.Logger

func InitLog(level zerolog.Level) *zerolog.Logger {
	zerolog.SetGlobalLevel(level)

	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()

	return &log
}
