package named

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgermode "github.com/ralvarezdev/go-logger/mode"
)

type (
	// DefaultLogger is the default mode logger that prefixes messages with a header.
	DefaultLogger struct {
		header string
		logger gologgermode.Logger
	}
)

// NewDefaultLogger creates a new DefaultLogger with the given header and logger.
//
// Parameters:
//
//	header - the prefix for all log messages.
//	logger - the underlying logger to delegate to.
//
// Returns:
//
//	A pointer to DefaultLogger and an error if logger is nil.
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

// Info logs an informational message with the DefaultLogger's header.
//
// Parameters:
//
//	subheader - a short description of the message.
//	details - optional additional details.
func (d DefaultLogger) Info(subheader string, details ...string) {
	d.logger.Info(d.header, subheader, &details)
}

// Error logs an error message with the DefaultLogger's header.
//
// Parameters:
//
//	subheader - a short description of the error.
//	errors - the error values to log.
func (d DefaultLogger) Error(subheader string, errors ...error) {
	d.logger.Error(d.header, subheader, &errors)
}

// Debug logs a debug message with the DefaultLogger's header.
//
// Parameters:
//
//	subheader - a short description of the message.
//	details - optional additional details.
func (d DefaultLogger) Debug(subheader string, details ...string) {
	d.logger.Debug(d.header, subheader, &details)
}

// Critical logs a critical message with the DefaultLogger's header.
//
// Parameters:
//
//	subheader - a short description of the message.
//	details - optional additional details.
func (d DefaultLogger) Critical(subheader string, details ...string) {
	d.logger.Critical(d.header, subheader, &details)
}

// Warning logs a warning message with the DefaultLogger's header.
//
// Parameters:
//
//	subheader - a short description of the message.
//	details - optional additional details.
func (d DefaultLogger) Warning(subheader string, details ...string) {
	d.logger.Warning(d.header, subheader, &details)
}
