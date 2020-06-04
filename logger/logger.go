package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Info(msg string) {
	log.Info().Msg(msg)
}

func Str(key string, value string) *zerolog.Event {
	return log.Info().Str(key, value)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}
