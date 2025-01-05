package mode

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gologger "github.com/ralvarezdev/go-logger"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
)

type (
	// Logger interface for the mode logger
	Logger interface {
		ShouldLog(status gologgerstatus.Status) bool
		RunIfShouldLog(status gologgerstatus.Status, fn func())
		gologger.Logger
	}

	// DefaultLogger is the default mode logger
	DefaultLogger struct {
		logger gologger.Logger
	}
)

// NewDefaultLogger creates a new default mode logger
func NewDefaultLogger(logger gologger.Logger) (
	*DefaultLogger,
	error,
) {
	// Check if the logger is nil
	if logger == nil {
		return nil, gologger.ErrNilLogger
	}

	return &DefaultLogger{logger: logger}, nil
}

// Log logs a message
func (d *DefaultLogger) Log(message *gologger.Message) {
	// Check if the message is nil
	if message == nil {
		return
	}

	d.RunIfShouldLog(
		message.Status(), func() {
			d.logger.Log(message)
		},
	)
}

// ShouldLog checks if the log should be logged
func (d *DefaultLogger) ShouldLog(status gologgerstatus.Status) bool {
	return LogModeMap[goflagsmode.ModeFlag.Mode()][status]
}

// RunIfShouldLog runs the function if the log should be logged
func (d *DefaultLogger) RunIfShouldLog(
	status gologgerstatus.Status,
	fn func(),
) {
	if d.ShouldLog(status) {
		fn()
	}
}

// Info logs an info message
func (d *DefaultLogger) Info(header, subheader string, details *[]string) {
	d.RunIfShouldLog(
		gologgerstatus.Info, func() {
			d.logger.Info(
				header,
				subheader,
				details,
			)
		},
	)
}

// Error logs an error message
func (d *DefaultLogger) Error(header, subheader string, errors *[]error) {
	d.RunIfShouldLog(
		gologgerstatus.Error, func() {
			d.logger.Error(
				header,
				subheader,
				errors,
			)
		},
	)
}

// Debug logs a debug message
func (d *DefaultLogger) Debug(header, subheader string, details *[]string) {
	d.RunIfShouldLog(
		gologgerstatus.Debug, func() {
			d.logger.Debug(
				header,
				subheader,
				details,
			)
		},
	)
}

// Critical logs a critical message
func (d *DefaultLogger) Critical(header, subheader string, details *[]string) {
	d.RunIfShouldLog(
		gologgerstatus.Critical, func() {
			d.logger.Critical(
				header,
				subheader,
				details,
			)
		},
	)
}

// Warning logs a warning message
func (d *DefaultLogger) Warning(header, subheader string, details *[]string) {
	d.RunIfShouldLog(
		gologgerstatus.Warning, func() {
			d.logger.Warning(
				header,
				subheader,
				details,
			)
		},
	)
}
