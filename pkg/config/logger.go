package config

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const FILE = "file"
const STDERR = "stderr"
const STDOUT = "stdout"
const HUMAN = "human"
const HUMANERR = "humanerr"

type loggerTarget struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}

func initLogger() {
	var loggers []zerolog.Logger
	for _, writer := range getLoggerWriters() {
		logger := zerolog.
			New(writer).
			With().
			Str("env", viper.GetString("env")).
			Str("role", viper.GetString("role")).
			Timestamp().
			Caller().
			Logger()
		loggers = append(loggers, logger)
	}

	log.SetFlags(0)
	if len(loggers) == 0 {
		log.Fatal("No zerolog loggers where instantiated!")
	}

	log.SetOutput(loggers[0])
	initLoggerGlobals(loggers[0])
}

func initLoggerGlobals(logger zerolog.Logger) {
	errorlevel, errGL := zerolog.ParseLevel(viper.GetString("logger.errorlevel"))
	if errGL != nil {
		log.Fatalf("Unable to parse global log level, %v", errGL)
	}
	zerolog.SetGlobalLevel(errorlevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zlog.Logger = logger
}

func getLoggerWriters() []io.Writer {
	var writers []io.Writer

	var targets []*loggerTarget
	err := viper.UnmarshalKey("logger.targets", &targets)
	if err != nil {
		log.Fatalf("Unable to decode 'logger.targets.default' into struct, %v", err)
	}

	for _, target := range targets {
		switch target.Type {
		case FILE:
			filePath := target.Path
			file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			// defer file.Close()
			writers = append(writers, file)
		case STDERR:
			writers = append(writers, os.Stderr)
		case STDOUT:
			writers = append(writers, os.Stdout)
		case HUMAN:
			writers = append(writers, zerolog.ConsoleWriter{Out: os.Stdout})
		case HUMANERR:
			writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
		default:
			panic(fmt.Sprintf("Unknown logger: %v", target))
		}
	}

	return writers
}
