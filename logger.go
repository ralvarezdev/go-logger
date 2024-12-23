package go_logger

import (
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
	gologgerstrings "github.com/ralvarezdev/go-logger/strings"
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
		status  gologgerstatus.Status
		format  *LogMessageFormat
	}

	LogMessageFormat struct {
		StatusSeparator  *gologgerstrings.ContentSeparator
		DetailsSeparator *gologgerstrings.MultilineSeparator
	}

	// LogError struct
	LogError struct {
		title  string
		errors []error
		format *LogErrorFormat
	}

	LogErrorFormat struct {
		StatusSeparator *gologgerstrings.ContentSeparator
		ErrorsSeparator *gologgerstrings.MultilineSeparator
	}

	// Logger is an interface for logging messages
	Logger interface {
		LogMessage(logMessage *LogMessage)
		LogError(logError *LogError)
	}

	// LoggerFormat struct
	LoggerFormat struct {
		NameSeparator    *gologgerstrings.ContentSeparator
		MessageSeparator gologgerstrings.Separator
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
	StatusSeparator:  gologgerstrings.NewRepeatedContentSeparator(gologgerstrings.SpaceSeparator),
	DetailsSeparator: gologgerstrings.NewMultilineSeparator(gologgerstrings.NewLineSeparator, 1),
}

// DefaultLogErrorFormat is the default log error format
var DefaultLogErrorFormat = LogErrorFormat{
	StatusSeparator: gologgerstrings.NewRepeatedContentSeparator(gologgerstrings.SpaceSeparator),
	ErrorsSeparator: gologgerstrings.NewMultilineSeparator(gologgerstrings.NewLineSeparator, 1),
}

// DefaultLoggerFormat is the default logger format
var DefaultLoggerFormat = LoggerFormat{
	NameSeparator:    gologgerstrings.NewRepeatedContentSeparator(gologgerstrings.SpaceSeparator),
	MessageSeparator: gologgerstrings.SpaceSeparator,
}

// NewLogMessageFormat creates a new log message format
func NewLogMessageFormat(
	statusSeparator *gologgerstrings.ContentSeparator,
	detailsSeparator *gologgerstrings.MultilineSeparator,
) *LogMessageFormat {
	return &LogMessageFormat{
		StatusSeparator:  statusSeparator,
		DetailsSeparator: detailsSeparator,
	}
}

// CopyLogMessageFormat creates a copy of a log message format
func CopyLogMessageFormat(format *LogMessageFormat) *LogMessageFormat {
	return &LogMessageFormat{
		StatusSeparator:  format.StatusSeparator,
		DetailsSeparator: format.DetailsSeparator,
	}
}

// NewLogErrorFormat creates a new log error format
func NewLogErrorFormat(
	statusSeparator *gologgerstrings.ContentSeparator,
	errorsSeparator *gologgerstrings.MultilineSeparator,
) *LogErrorFormat {
	return &LogErrorFormat{
		StatusSeparator: statusSeparator,
		ErrorsSeparator: errorsSeparator,
	}
}

// CopyLogErrorFormat creates a copy of a log error format
func CopyLogErrorFormat(format *LogErrorFormat) *LogErrorFormat {
	return &LogErrorFormat{
		StatusSeparator: format.StatusSeparator,
		ErrorsSeparator: format.ErrorsSeparator,
	}
}

// NewLoggerFormat creates a new logger format
func NewLoggerFormat(
	nameSeparator *gologgerstrings.ContentSeparator, messageSeparator gologgerstrings.Separator,
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
func NewLogMessage(
	title string,
	status gologgerstatus.Status,
	format *LogMessageFormat,
	details ...string,
) *LogMessage {
	// Check if the format is nil
	if format == nil {
		format = &DefaultLogMessageFormat
	}

	return &LogMessage{title: title, status: status, details: details, format: CopyLogMessageFormat(format)}
}

// FormatDetails gets the formatted details
func (l *LogMessage) FormatDetails() string {
	return gologgerstrings.FormatStringArray(l.format.DetailsSeparator, &l.details)
}

// String gets the string representation of a log message
func (l *LogMessage) String() string {
	var formattedLog []string

	// Format status
	if l.status != gologgerstatus.StatusNone {
		formattedLog = append(formattedLog, gologgerstrings.FormatStatus(l.status, l.format.StatusSeparator))
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
	return gologgerstrings.FormatErrorArray(l.format.ErrorsSeparator, &l.errors)
}

// String gets the string representation of a log error
func (l *LogError) String() string {
	var formattedLog []string

	// Format status
	formattedLog = append(
		formattedLog,
		gologgerstrings.FormatStatus(gologgerstatus.StatusFailed, l.format.StatusSeparator),
	)

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
		formattedName: gologgerstrings.AddBrackets(name, format.NameSeparator),
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
