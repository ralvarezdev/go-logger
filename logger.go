package go_logger

import (
	gologgerseparator "github.com/ralvarezdev/go-logger/separator"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
	gologgerstrings "github.com/ralvarezdev/go-logger/strings"
	"log"
	"strings"
)

var (
	// StatusSeparator is the status separator
	StatusSeparator = gologgerseparator.NewRepeatedContent(gologgerseparator.Space)

	// ErrorsSeparator is the errors separator
	ErrorsSeparator = gologgerseparator.NewMultiline(
		gologgerseparator.Space,
		gologgerseparator.NewLine,
		1,
	)

	// DetailsSeparator is the details separator
	DetailsSeparator = gologgerseparator.NewMultiline(
		gologgerseparator.Space,
		gologgerseparator.NewLine,
		1,
	)

	// NameSeparator is the name separator
	NameSeparator = gologgerseparator.NewRepeatedContent(gologgerseparator.Space)

	// MessageSeparator is the message separator
	MessageSeparator = gologgerseparator.Space
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
	}

	// LogError struct
	LogError struct {
		title  string
		errors []error
	}

	// Logger is an interface for logging messages
	Logger interface {
		LogMessage(logMessage *LogMessage)
		LogError(logError *LogError)
	}

	// DefaultLogger is a logger that logs messages
	DefaultLogger struct {
		name          string
		formattedName string
	}
)

// NewLogMessage creates a new log message
func NewLogMessage(
	title string,
	status gologgerstatus.Status,
	details ...string,
) *LogMessage {
	return &LogMessage{
		title:   title,
		status:  status,
		details: details,
	}
}

// FormatDetails gets the formatted details
func (l *LogMessage) FormatDetails() string {
	return gologgerstrings.FormatStringArray(
		DetailsSeparator,
		&l.details,
	)
}

// String gets the string representation of a log message
func (l *LogMessage) String() string {
	var formattedLog []string

	// Format status
	if l.status != gologgerstatus.None {
		formattedLog = append(
			formattedLog,
			gologgerstrings.FormatStatus(StatusSeparator, l.status),
		)
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
func NewLogError(
	title string,
	errors ...error,
) *LogError {
	return &LogError{
		title:  title,
		errors: errors,
	}
}

// FormatErrors gets the formatted errors
func (l *LogError) FormatErrors() string {
	return gologgerstrings.FormatErrorArray(ErrorsSeparator, &l.errors)
}

// String gets the string representation of a log error
func (l *LogError) String() string {
	var formattedLog []string

	// Format status
	formattedLog = append(
		formattedLog,
		gologgerstrings.FormatStatus(
			StatusSeparator,
			gologgerstatus.Failed,
		),
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
func NewDefaultLogger(name string) *DefaultLogger {
	return &DefaultLogger{
		name:          name,
		formattedName: gologgerstrings.AddBrackets(NameSeparator, name),
	}
}

// FormatLogMessage formats a log message
func (d *DefaultLogger) FormatLogMessage(logMessage *LogMessage) string {
	// Check if the log message is nil
	if logMessage == nil {
		return d.formattedName
	}

	return strings.Join(
		[]string{d.formattedName, logMessage.String()},
		string(MessageSeparator),
	)
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
		}, string(MessageSeparator),
	)
}

// LogError logs an error
func (d *DefaultLogger) LogError(logError *LogError) {
	log.Println(d.FormatLogError(logError))
}
