package log

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var logLevelsMap = map[string]zerolog.Level{
	"trace":    zerolog.TraceLevel,
	"debug":    zerolog.DebugLevel,
	"info":     zerolog.InfoLevel,
	"warn":     zerolog.WarnLevel,
	"error":    zerolog.ErrorLevel,
	"fatal":    zerolog.FatalLevel,
	"panic":    zerolog.PanicLevel,
	"disabled": zerolog.Disabled,
}

func InitLog(logLevel string) *zerolog.Logger {
	var log zerolog.Logger

	output := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}

	if l, ok := logLevelsMap[logLevel]; ok {
		log = zerolog.New(output).With().Timestamp().Logger().Level(l)
	} else {
		cobra.CheckErr(fmt.Errorf("Loglevel '%v' is incorrect. Possible values is: race, debug, info, warn, error, fatal, panic, disable.", logLevel))
	}

	return &log
}
