package logger

import (
	"io"
	"strings"

	"github.com/rs/zerolog"
)

type Logger interface {
	Panic(string)
	Fatal(string)
	Error(string)
	Warning(string)
	Info(string)
	Debug(string)
	With(key, value string) Logger
	AddError(err error) Logger
}

type logger struct {
	log zerolog.Logger
}

const (
	Noop    = "noop"
	Panic   = "panic"
	Fatal   = "fatal"
	Error   = "error"
	Warning = "warn"
	Info    = "info"
	Debug   = "debug"
)

func New(output io.Writer, level string) (Logger, error) {
	logLevel, err := parseLogLevel(level)
	if err != nil {
		return nil, err
	}

	zl := zerolog.
		New(output).
		Level(logLevel).
		With().
		Timestamp().
		Logger()

	return &logger{
		log: zl,
	}, nil
}

func (l logger) Panic(msg string) {
	l.log.Panic().Msg(msg)
}

func (l logger) Fatal(msg string) {
	l.log.Fatal().Msg(msg)
}

func (l logger) Error(msg string) {
	l.log.Error().Msg(msg)
}

func (l logger) Warning(msg string) {
	l.log.Warn().Msg(msg)
}

func (l logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l logger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l logger) With(key, value string) Logger {
	return logger{
		log: l.log.With().Str(key, value).Logger(),
	}
}

func (l logger) AddError(err error) Logger {
	return logger{
		log: l.log.With().Err(err).Logger(),
	}
}

func parseLogLevel(level string) (zerolog.Level, error) {
	switch strings.ToLower(level) {
	case Noop:
		return zerolog.Disabled, nil
	case Panic:
		return zerolog.PanicLevel, nil
	case Fatal:
		return zerolog.FatalLevel, nil
	case Error:
		return zerolog.ErrorLevel, nil
	case Warning:
		return zerolog.WarnLevel, nil
	case Info:
		return zerolog.InfoLevel, nil
	case Debug:
		return zerolog.DebugLevel, nil
	}

	return zerolog.Disabled, UnknownLogLevel(level)
}
