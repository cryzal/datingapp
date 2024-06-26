package log

import (
	"context"
	LogService "datingapp/shared/utils/logger"
	"io"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	logger = LogService.GetConfig()
}

func Output(w io.Writer) zerolog.Logger {
	return logger.Output(w)
}

func With() zerolog.Context {
	return logger.With()
}

func Level(level zerolog.Level) zerolog.Logger {
	return logger.Level(level)
}

func Sample(s zerolog.Sampler) zerolog.Logger {
	return logger.Sample(s)
}

func Hook(h zerolog.Hook) zerolog.Logger {
	return logger.Hook(h)
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Fatal() *zerolog.Event {
	return logger.Fatal()
}

func Panic() *zerolog.Event {
	return logger.Panic()
}

func WithLevel(level zerolog.Level) *zerolog.Event {
	return logger.WithLevel(level)
}

func Log() *zerolog.Event {
	return logger.Log()
}

func Print(v ...interface{}) {
	logger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
