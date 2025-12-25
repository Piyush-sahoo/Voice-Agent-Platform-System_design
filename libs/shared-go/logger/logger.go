package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// Logger is the global logger instance
var Logger zerolog.Logger

func init() {
	Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

// WithService returns a logger with service name attached
func WithService(service string) zerolog.Logger {
	return Logger.With().Str("service", service).Logger()
}

// WithCallID returns a logger with call ID attached
func WithCallID(callID string) zerolog.Logger {
	return Logger.With().Str("call_id", callID).Logger()
}
