package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var (
	logger *zerolog.Logger
)

func Init(releaseMode string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if releaseMode == gin.ReleaseMode {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)

		logRotator := &lumberjack.Logger{
			Filename:   "/tmp/log/structure.log",
			MaxBackups: 3,
			MaxSize:    500,
			MaxAge:     3,
			Compress:   true,
		}

		l := zerolog.New(logRotator).With().Timestamp().CallerWithSkipFrameCount(3).Logger()
		logger = &l
		return
	}

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%+v", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%+v:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%+v", i)
	}

	l := zerolog.New(output).With().Timestamp().CallerWithSkipFrameCount(3).Logger()
	logger = &l
}

func Logger() *zerolog.Logger {
	return logger
}

func Error(err string) {
	logger.Error().Msg(err)
}

func Errorf(format string, v ...interface{}) {
	logger.Error().Msgf(format, v...)
}

func Debug(msg string) {
	logger.Debug().Msg(msg)
}

func Debugf(format string, v ...interface{}) {
	logger.Debug().Msgf(format, v...)
}

func Info(msg string) {
	logger.Info().Msg(msg)
}

func Infof(format string, v ...interface{}) {
	logger.Info().Msgf(format, v...)
}

func Warn(msg string) {
	logger.Warn().Msg(msg)
}

func Warnf(format string, v ...interface{}) {
	logger.Warn().Msgf(format, v...)
}

func Fatal(msg string) {
	logger.Fatal().Msg(msg)
}
