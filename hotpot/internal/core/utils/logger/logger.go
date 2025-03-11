// Package logger provides a structured and configurable logging utility
// for Go applications, leveraging Go's slog package with JSON output
// and deduplication support to avoid duplicate log keys.
package logger

import (
	"fmt"
	"log/slog"
	"os"

	slogdd "github.com/veqryn/slog-dedup"
)

// Config represents the configuration for the logger.
// It allows setting the log level to control verbosity.
type Config struct {
	Level string // Log level: "debug", "info", "warn", or "error".
}

// DefaultConfig returns the default logger configuration.
// By default, the log level is set to "info".
func DefaultConfig() *Config {
	return &Config{
		Level: "info",
	}
}

// Validate checks if the configuration is valid.
// It ensures the log level is one of the supported values.
func (c *Config) Validate() error {
	switch c.Level {
	case "debug", "info", "warn", "error":
	default:
		return fmt.Errorf("unknown log level: %s", c.Level)
	}
	return nil
}

// New creates a new slog.Logger instance based on the provided configuration.
// The logger outputs JSON-formatted logs and deduplicates log keys to ensure
// compatibility with log-parsing tools such as Elasticsearch.
//
// Arguments:
//
//	config - A validated Config instance that specifies the log level.
//
// Returns:
//
//	A pointer to an instance of slog.Logger configured with JSON output
//	and deduplication.
func New(config *Config) *slog.Logger {
	var slogLvl slog.Level
	_ = slogLvl.UnmarshalText([]byte(config.Level))

	return slog.New(
		// Using deduplication to handle potential duplicate log keys,
		// particularly important when exporting logs to JSON-parsing tools.
		// Reference: https://github.com/golang/go/issues/59365
		slogdd.NewIncrementHandler(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slogLvl},
			),
			nil,
		))
}
