package go_logger

import (
	"github.com/ralvarezdev/go-logger/status"
	loggerstrings "github.com/ralvarezdev/go-logger/strings"
	"log"
	"strings"
)

type (
	// Log interface
	Log interface {
		String() string
	}

	// LogMessage struct
	LogMessage struct {
		title   string
		details []string
		status  status.Status
		format  *LogMessageFormat
	}

	LogMessageFormat struct {
		StatusSeparator       loggerstrings.Separator
		DetailsOuterSeparator loggerstrings.Separator
		DetailsInnerSeparator loggerstrings.Separator
	}

	// LogError struct
	LogError struct {
		title  string
		errors []error
		format *LogErrorFormat
	}

	LogErrorFormat struct {
		StatusSeparator      loggerstrings.Separator
		ErrorsOuterSeparator loggerstrings.Separator
		ErrorsInnerSeparator loggerstrings.Separator
	}

	// Logger is an interface for logging messages
	Logger interface {
		LogMessage(logMessage *LogMessage)
		LogError(logError *LogError)
	}

	// LoggerFormat struct
	LoggerFormat struct {
		NameSeparator    loggerstrings.Separator
		MessageSeparator loggerstrings.Separator
	}

	// DefaultLogger is a logger that logs messages
	DefaultLogger struct {
		name          string
		formattedName string
		format        *LoggerFormat
	}
)

// DefaultLogMessageFormat is the default log message format
var DefaultLogMessageFormat = LogMessageFormat{
	StatusSeparator:       loggerstrings.SpaceSeparator,
	DetailsOuterSeparator: loggerstrings.NewLineSeparator,
	DetailsInnerSeparator: loggerstrings.TabSeparator,
}

// DefaultLogErrorFormat is the default log error format
var DefaultLogErrorFormat = LogErrorFormat{
	StatusSeparator:      loggerstrings.SpaceSeparator,
	ErrorsOuterSeparator: loggerstrings.NewLineSeparator,
	ErrorsInnerSeparator: loggerstrings.TabSeparator,
}

// DefaultLoggerFormat is the default logger format
var DefaultLoggerFormat = LoggerFormat{
	NameSeparator:    loggerstrings.SpaceSeparator,
	MessageSeparator: loggerstrings.SpaceSeparator,
}

// NewLogMessageFormat creates a new log message format
func NewLogMessageFormat(
	statusSeparator, detailsOuterSeparator, detailsInnerSeparator loggerstrings.Separator,
) *LogMessageFormat {
	return &LogMessageFormat{
		StatusSeparator:       statusSeparator,
		DetailsOuterSeparator: detailsOuterSeparator,
		DetailsInnerSeparator: detailsInnerSeparator,
	}
}

// CopyLogMessageFormat creates a copy of a log message format
func CopyLogMessageFormat(format *LogMessageFormat) *LogMessageFormat {
	return &LogMessageFormat{
		StatusSeparator:       format.StatusSeparator,
		DetailsOuterSeparator: format.DetailsOuterSeparator,
		DetailsInnerSeparator: format.DetailsInnerSeparator,
	}
}

// NewLogErrorFormat creates a new log error format
func NewLogErrorFormat(
	statusSeparator, errorsOuterSeparator, errorsInnerSeparator loggerstrings.Separator,
) *LogErrorFormat {
	return &LogErrorFormat{
		StatusSeparator:      statusSeparator,
		ErrorsOuterSeparator: errorsOuterSeparator,
		ErrorsInnerSeparator: errorsInnerSeparator,
	}
}

// CopyLogErrorFormat creates a copy of a log error format
func CopyLogErrorFormat(format *LogErrorFormat) *LogErrorFormat {
	return &LogErrorFormat{
		StatusSeparator:      format.StatusSeparator,
		ErrorsOuterSeparator: format.ErrorsOuterSeparator,
		ErrorsInnerSeparator: format.ErrorsInnerSeparator,
	}
}

// NewLoggerFormat creates a new logger format
func NewLoggerFormat(
	nameSeparator, messageSeparator loggerstrings.Separator,
) *LoggerFormat {
	return &LoggerFormat{
		NameSeparator:    nameSeparator,
		MessageSeparator: messageSeparator,
	}
}

// CopyLoggerFormat creates a copy of a logger format
func CopyLoggerFormat(format *LoggerFormat) *LoggerFormat {
	return &LoggerFormat{
		NameSeparator:    format.NameSeparator,
		MessageSeparator: format.MessageSeparator,
	}
}

// NewLogMessage creates a new log message
func NewLogMessage(title string, status status.Status, format *LogMessageFormat, details ...string) *LogMessage {
	// Check if the format is nil
	if format == nil {
		format = &DefaultLogMessageFormat
	}

	return &LogMessage{title: title, status: status, details: details, format: CopyLogMessageFormat(format)}
}

// FormatDetails gets the formatted details
func (l *LogMessage) FormatDetails() string {
	return loggerstrings.FormatStringArray(l.format.DetailsOuterSeparator, l.format.DetailsInnerSeparator, &l.details)
}

// String gets the string representation of a log message
func (l *LogMessage) String() string {
	var formattedLog []string

	// Format status
	if l.status != status.StatusNone {
		formattedLog = append(formattedLog, loggerstrings.FormatStatus(l.status, l.format.StatusSeparator))
	}

	// Add title
	if l.title != "" {
		formattedLog = append(formattedLog, l.title)
	}

	// Add formatted details
	if len(l.details) > 0 {
		formattedLog = append(formattedLog, l.FormatDetails())
	}

	return strings.Join(formattedLog, " ")
}

// NewLogError creates a new log error
func NewLogError(title string, format *LogErrorFormat, errors ...error) *LogError {
	// Check if the format is nil
	if format == nil {
		format = &DefaultLogErrorFormat
	}

	return &LogError{title: title, errors: errors, format: CopyLogErrorFormat(format)}
}

// FormatErrors gets the formatted errors
func (l *LogError) FormatErrors() string {
	return loggerstrings.FormatErrorArray(l.format.ErrorsOuterSeparator, l.format.ErrorsInnerSeparator, &l.errors)
}

// String gets the string representation of a log error
func (l *LogError) String() string {
	var formattedLog []string

	// Format status
	formattedLog = append(formattedLog, loggerstrings.FormatStatus(status.StatusFailed, l.format.StatusSeparator))

	// Add message
	if l.title != "" {
		formattedLog = append(formattedLog, l.title)
	}

	// Add formatted errors
	if len(l.errors) > 0 {
		formattedLog = append(formattedLog, l.FormatErrors())
	}

	return strings.Join(formattedLog, " ")
}

// NewDefaultLogger creates a new logger
func NewDefaultLogger(name string, format *LoggerFormat) *DefaultLogger {
	// Check if the format is nil
	if format == nil {
		format = &DefaultLoggerFormat
	}

	return &DefaultLogger{
		name:          name,
		formattedName: loggerstrings.AddBrackets(name, format.NameSeparator),
		format:        CopyLoggerFormat(format),
	}
}

// FormatLogMessage formats a log message
func (d *DefaultLogger) FormatLogMessage(logMessage *LogMessage) string {
	// Check if the log message is nil
	if logMessage == nil {
		return d.formattedName
	}

	return strings.Join([]string{d.formattedName, logMessage.String()}, string(d.format.MessageSeparator))
}

// LogMessage logs a message
func (d *DefaultLogger) LogMessage(logMessage *LogMessage) {
	log.Println(d.FormatLogMessage(logMessage))
}

// FormatLogError formats a log error
func (d *DefaultLogger) FormatLogError(logError *LogError) string {
	// Check if the log error is nil
	if logError == nil {
		return d.formattedName
	}

	return strings.Join(
		[]string{
			d.formattedName,
			logError.String(),
		}, string(d.format.MessageSeparator),
	)
}

// LogError logs an error
func (d *DefaultLogger) LogError(logError *LogError) {
	log.Println(d.FormatLogError(logError))
}
