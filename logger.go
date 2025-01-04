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
		subheader string
		details   []string
		status    gologgerstatus.Status
	}

	// LogError struct
	LogError struct {
		subheader string
		errors    []error
	}

	// Logger is an interface for logging messages
	Logger interface {
		LogMessage(header string, logMessage *LogMessage)
		LogError(header string, logError *LogError)
	}

	// SubLogger is an interface for logging messages with a name
	SubLogger interface {
		LogMessage(logMessage *LogMessage)
		LogError(logError *LogError)
	}

	// DefaultLogger is a logger that logs messages
	DefaultLogger struct{}

	// DefaultSubLogger is a logger that logs messages with a name
	DefaultSubLogger struct {
		name          string
		formattedName string
		logger        Logger
	}
)

// NewLogMessage creates a new log message
func NewLogMessage(
	subheader string,
	status gologgerstatus.Status,
	details ...string,
) *LogMessage {
	return &LogMessage{
		subheader: subheader,
		status:    status,
		details:   details,
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

	// Add subheader
	if l.subheader != "" {
		formattedLog = append(formattedLog, l.subheader)
	}

	// Add formatted details
	if len(l.details) > 0 {
		formattedLog = append(formattedLog, l.FormatDetails())
	}

	return strings.Join(formattedLog, " ")
}

// NewLogError creates a new log error
func NewLogError(
	subheader string,
	errors ...error,
) *LogError {
	return &LogError{
		subheader: subheader,
		errors:    errors,
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

	// Add subheader
	if l.subheader != "" {
		formattedLog = append(formattedLog, l.subheader)
	}

	// Add formatted errors
	if len(l.errors) > 0 {
		formattedLog = append(formattedLog, l.FormatErrors())
	}

	return strings.Join(formattedLog, " ")
}

// NewDefaultLogger creates a new logger
func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

// FormatLogMessage formats a log message
func (d *DefaultLogger) FormatLogMessage(
	header string,
	logMessage *LogMessage,
) string {
	return strings.Join(
		[]string{header, logMessage.String()},
		string(MessageSeparator),
	)
}

// LogMessage logs a message
func (d *DefaultLogger) LogMessage(header string, logMessage *LogMessage) {
	log.Println(d.FormatLogMessage(header, logMessage))
}

// FormatLogError formats a log error
func (d *DefaultLogger) FormatLogError(
	header string,
	logError *LogError,
) string {
	return strings.Join(
		[]string{
			header,
			logError.String(),
		}, string(MessageSeparator),
	)
}

// LogError logs an error
func (d *DefaultLogger) LogError(header string, logError *LogError) {
	log.Println(d.FormatLogError(header, logError))
}

// NewDefaultSubLogger creates a new sub logger
func NewDefaultSubLogger(name string, logger Logger) (
	*DefaultSubLogger,
	error,
) {
	// Check if the logger is nil
	if logger == nil {
		return nil, ErrNilLogger
	}

	return &DefaultSubLogger{
		name:          name,
		formattedName: gologgerstrings.AddBrackets(NameSeparator, name),
		logger:        logger,
	}, nil
}

// LogMessage logs a message
func (d *DefaultSubLogger) LogMessage(logMessage *LogMessage) {
	d.logger.LogMessage(d.formattedName, logMessage)
}

// LogError logs an error
func (d *DefaultSubLogger) LogError(logError *LogError) {
	d.logger.LogError(d.formattedName, logError)
}
