package go_logger

import (
	gologgerseparator "github.com/ralvarezdev/go-logger/separator"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
	gologgerstrings "github.com/ralvarezdev/go-logger/strings"
	"log"
	"strings"
)

var (
	// HeaderSeparator is the header separator
	HeaderSeparator = gologgerseparator.NewRepeatedContent(gologgerseparator.Space)

	// StatusSeparator is the status separator
	StatusSeparator = gologgerseparator.NewRepeatedContent(gologgerseparator.Space)

	// DescriptionSeparator is the description separator
	DescriptionSeparator = gologgerseparator.NewMultiline(
		gologgerseparator.Space,
		gologgerseparator.NewLine,
		1,
	)

	// MessageSeparator is the message separator
	MessageSeparator = gologgerseparator.Space

	// AddCharactersFn is the add characters function
	AddCharactersFn = gologgerstrings.AddBrackets
)

type (
	// Message struct
	Message struct {
		header      string
		subheader   string
		description *[]string
		status      gologgerstatus.Status
	}

	// Logger is an interface for logging messages
	Logger interface {
		Log(message *Message)
		Info(header, subheader string, details ...string)
		Error(header, subheader string, errors ...error)
		Debug(header, subheader string, details ...string)
		Critical(header, subheader string, details ...string)
		Warning(header, subheader string, details ...string)
	}

	// DefaultLogger is a logger that logs messages
	DefaultLogger struct{}
)

// NewMessage creates a new message
func NewMessage(
	header, subheader string,
	status gologgerstatus.Status,
	description *[]string,
) *Message {
	return &Message{
		header:      header,
		subheader:   subheader,
		status:      status,
		description: description,
	}
}

// String gets the string representation of a message
func (m *Message) String() string {
	var formattedMessage []string

	// Add header
	if m.header != "" {
		formattedMessage = append(
			formattedMessage,
			gologgerstrings.FormatString(
				HeaderSeparator,
				m.header,
				AddCharactersFn,
			),
		)
	}

	// Format status
	formattedMessage = append(
		formattedMessage,
		gologgerstrings.FormatStatus(
			StatusSeparator,
			m.status,
			AddCharactersFn,
		),
	)

	// Add subheader
	if m.subheader != "" {
		formattedMessage = append(formattedMessage, m.subheader)
	}

	// Add formatted description
	if m.description != nil && len(*m.description) > 0 {
		formattedMessage = append(
			formattedMessage, gologgerstrings.FormatStringArray(
				DescriptionSeparator,
				m.description,
				AddCharactersFn,
			),
		)
	}

	return strings.Join(formattedMessage, string(MessageSeparator))
}

// NewDefaultLogger creates a new logger
func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

// Log logs a message
func (d *DefaultLogger) Log(message *Message) {
	log.Println(message.String())
}

// BuildAndLog builds a message and logs it
func (d *DefaultLogger) BuildAndLog(
	header, subheader string,
	details *[]string,
	status gologgerstatus.Status,
) {
	// Create a new message and log it
	message := NewMessage(
		header,
		subheader,
		status,
		details,
	)
	d.Log(message)
}

// Info logs an info message
func (d *DefaultLogger) Info(header, subheader string, details ...string) {
	d.BuildAndLog(
		header,
		subheader,
		&details,
		gologgerstatus.Info,
	)
}

// Error logs an error message
func (d *DefaultLogger) Error(header, subheader string, errors ...error) {
	// Map errors to a string array
	mappedErrors := gologgerstrings.MapErrorArrayToStringArray(&errors)
	d.BuildAndLog(
		header,
		subheader,
		mappedErrors,
		gologgerstatus.Error,
	)
}

// Debug logs a debug message
func (d *DefaultLogger) Debug(header, subheader string, details ...string) {
	d.BuildAndLog(
		header,
		subheader,
		&details,
		gologgerstatus.Debug,
	)
}

// Critical logs a critical message
func (d *DefaultLogger) Critical(header, subheader string, details ...string) {
	d.BuildAndLog(
		header,
		subheader,
		&details,
		gologgerstatus.Critical,
	)
}

// Warning logs a warning message
func (d *DefaultLogger) Warning(header, subheader string, details ...string) {
	d.BuildAndLog(
		header,
		subheader,
		&details,
		gologgerstatus.Warning,
	)
}
