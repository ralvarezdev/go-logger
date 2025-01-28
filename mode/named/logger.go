package named

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgermode "github.com/ralvarezdev/go-logger/mode"
)

type (
	// Logger interface for the mode logger
	Logger interface {
		Info(subheader string, details ...string)
		Error(subheader string, errors ...error)
		Debug(subheader string, details ...string)
		Critical(subheader string, details ...string)
		Warning(subheader string, details ...string)
	}

	// DefaultLogger is the default mode logger
	DefaultLogger struct {
		header string
		logger gologgermode.Logger
	}
)

// NewDefaultLogger creates a new default mode logger
func NewDefaultLogger(header string, logger gologgermode.Logger) (
	*DefaultLogger,
	error,
) {
	// Check if the logger is nil
	if logger == nil {
		return nil, gologger.ErrNilLogger
	}

	return &DefaultLogger{header, logger}, nil
}

// Info logs an info message
func (d *DefaultLogger) Info(subheader string, details ...string) {
	d.logger.Info(d.header, subheader, &details)
}

// Error logs an error message
func (d *DefaultLogger) Error(subheader string, errors ...error) {
	d.logger.Error(d.header, subheader, &errors)
}

// Debug logs a debug message
func (d *DefaultLogger) Debug(subheader string, details ...string) {
	d.logger.Debug(d.header, subheader, &details)
}

// Critical logs a critical message
func (d *DefaultLogger) Critical(subheader string, details ...string) {
	d.logger.Critical(d.header, subheader, &details)
}

// Warning logs a warning message
func (d *DefaultLogger) Warning(subheader string, details ...string) {
	d.logger.Warning(d.header, subheader, &details)
}
