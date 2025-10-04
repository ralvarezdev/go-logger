package go_logger

import (
	"log"
	"strings"

	gologgerstatus "github.com/ralvarezdev/go-logger/status"
	gostringsaddformat "github.com/ralvarezdev/go-strings/add/format"
	gostringsconvert "github.com/ralvarezdev/go-strings/convert"
)

type (
	// Message struct
	Message struct {
		header      string
		subheader   string
		description *[]string
		status      gologgerstatus.Status
	}

	// DefaultLogger is a logger that logs messages
	DefaultLogger struct{}
)

// NewMessage creates a new message
//
// Parameters:
//
//   - header: the header of the message
//   - subheader: the subheader of the message
//   - description: the description of the message
//   - status: the status of the message
//
// Returns:
//
//   - *Message: the new message
func NewMessage(
	header, subheader string,
	description *[]string,
	status gologgerstatus.Status,
) *Message {
	return &Message{
		header,
		subheader,
		description,
		status,
	}
}

// Status returns the status of a message
//
// Returns:
//
//   - gologgerstatus.Status: the status of the message
func (m Message) Status() gologgerstatus.Status {
	return m.status
}

// String gets the string representation of a message
//
// Returns:
//
//   - string: the string representation of the message
func (m Message) String() string {
	var formattedMessage []string

	// Add header
	if m.header != "" {
		formattedMessage = append(
			formattedMessage,
			AddCharactersFn(
				HeaderSeparator,
				m.header,
			),
		)
	}

	// Format status
	formattedMessage = append(
		formattedMessage,
		m.status.Format(
			StatusSeparator,
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
			formattedMessage, gostringsaddformat.StringArray(
				DescriptionSeparator,
				m.description,
				AddCharactersFn,
			),
		)
	}

	return strings.Join(formattedMessage, string(MessageSeparator))
}

// NewDefaultLogger creates a new logger
//
// Returns:
//
//   - *DefaultLogger: the new logger
func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

// Log logs a message
//
// Parameters:
//
//   - message: the message to log
func (d DefaultLogger) Log(message *Message) {
	log.Println(message.String())
}

// BuildAndLog builds a message and logs it
//
// Parameters:
//
//   - header: the header of the message
//   - subheader: the subheader of the message
//   - details: the details of the message
//   - status: the status of the message
func (d DefaultLogger) BuildAndLog(
	header, subheader string,
	details *[]string,
	status gologgerstatus.Status,
) {
	// Create a new message and log it
	message := NewMessage(
		header,
		subheader,
		details,
		status,
	)
	d.Log(message)
}

// Info logs an info message
//
// Parameters:
//
//   - header: the header of the message
//   - subheader: the subheader of the message
//   - details: the details of the message
func (d DefaultLogger) Info(header, subheader string, details *[]string) {
	d.BuildAndLog(
		header,
		subheader,
		details,
		gologgerstatus.Info,
	)
}

// Error logs an error message
//
// Parameters:
//
//   - header: the header of the message
//   - subheader: the subheader of the message
//   - errors: the errors of the message
func (d DefaultLogger) Error(header, subheader string, errors *[]error) {
	// Map errors to a string array
	mappedErrors := gostringsconvert.ErrorArrayToStringArray(errors)
	d.BuildAndLog(
		header,
		subheader,
		mappedErrors,
		gologgerstatus.Error,
	)
}

// Debug logs a debug message
//
// Parameters:
//
//   - header: the header of the message
//   - subheader: the subheader of the message
//   - details: the details of the message
func (d DefaultLogger) Debug(header, subheader string, details *[]string) {
	d.BuildAndLog(
		header,
		subheader,
		details,
		gologgerstatus.Debug,
	)
}

// Critical logs a critical message
//
// Parameters:
//
//   - header: the header of the message
//   - subheader: the subheader of the message
//   - details: the details of the message
func (d DefaultLogger) Critical(header, subheader string, details *[]string) {
	d.BuildAndLog(
		header,
		subheader,
		details,
		gologgerstatus.Critical,
	)
}

// Warning logs a warning message
//
// Parameters:
//
//   - header: the header of the message
//   - subheader: the subheader of the message
//   - details: the details of the message
func (d DefaultLogger) Warning(header, subheader string, details *[]string) {
	d.BuildAndLog(
		header,
		subheader,
		details,
		gologgerstatus.Warning,
	)
}
