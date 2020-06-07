package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

var (
	log zerolog.Logger
)

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	log = zerolog.New(output).With().Timestamp().Logger()
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Print(v ...interface{}) {
	data := ""
	for _, subData := range v {
		value, ok := subData.(string)
		if ok {
			data += value + " | "
		}
	}
	log.Print(data)
}

func Str(key string, value string) *zerolog.Event {
	return log.Info().Str(key, value)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Error(msg string) {
	log.Error().Msg(msg)
}
