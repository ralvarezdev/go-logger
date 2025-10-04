package mode

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gologger "github.com/ralvarezdev/go-logger"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
)

type (
	// DefaultLogger is the default mode logger
	DefaultLogger struct {
		logger   gologger.Logger
		flagMode *goflagsmode.Flag
	}
)

// NewDefaultLogger creates a new default mode logger
//
// Parameters:
//
//   - logger: the logger to use
//   - flagMode: the mode flag to use
func NewDefaultLogger(
	logger gologger.Logger,
	flagMode *goflagsmode.Flag,
) (
	*DefaultLogger,
	error,
) {
	// Check if the logger is nil
	if logger == nil {
		return nil, gologger.ErrNilLogger
	}

	return &DefaultLogger{logger, flagMode}, nil
}

// Log logs a message
//
// Parameters:
//
//   - message: the message to log
func (d DefaultLogger) Log(message *gologger.Message) {
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
//
// Parameters:
//
//   - status: the status of the log
//
// Returns:
//
//   - bool: true if the log should be logged, false otherwise
func (d DefaultLogger) ShouldLog(status gologgerstatus.Status) bool {
	if d.flagMode == nil {
		return true
	}
	if d.flagMode.IsDebug() {
		return true
	}
	return status != gologgerstatus.Debug
}

// RunIfShouldLog runs the function if the log should be logged
//
// Parameters:
//
//   - status: the status of the log
//   - fn: the function to run if the log should be logged
func (d DefaultLogger) RunIfShouldLog(
	status gologgerstatus.Status,
	fn func(),
) {
	if d.ShouldLog(status) {
		fn()
	}
}

// Info logs an info message
//
// Parameters:
//
//   - header: the header of the info message
//   - subheader: the subheader of the info message
//   - details: the details of the info message
func (d DefaultLogger) Info(header, subheader string, details *[]string) {
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
//
// Parameters:
//
//   - header: the header of the error message
//   - subheader: the subheader of the error message
//   - errors: the errors of the error message
func (d DefaultLogger) Error(header, subheader string, errors *[]error) {
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
func (d DefaultLogger) Debug(header, subheader string, details *[]string) {
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
//
// Parameters:
//
//   - header: the header of the critical message
//   - subheader: the subheader of the critical message
//   - details: the details of the critical message
func (d DefaultLogger) Critical(header, subheader string, details *[]string) {
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
//
// Parameters:
//
//   - header: the header of the warning message
//   - subheader: the subheader of the warning message
//   - details: the details of the warning message
func (d DefaultLogger) Warning(header, subheader string, details *[]string) {
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
