package logger

import (
	"log/slog"
	"os"
)

// getLogger returns the logger to use in the entire application.
func getLogger() *slog.Logger {
	return slog.Default()
}

// Info logs an info message using the configured logger.
func Info(msg string, args ...any) {
	getLogger().Info(msg, args...)
}

// Warn logs a warning message using the configured logger.
func Warn(msg string, args ...any) {
	getLogger().Warn(msg, args...)
}

// Error logs an error message using the configured logger.
func Error(msg string, args ...any) {
	getLogger().Error(msg, args...)
}

// FatalError is equivalent to Error() followed by a call to os.Exit(1)
func FatalError(msg string, args ...any) {
	Error(msg, args...)
	os.Exit(1)
}
