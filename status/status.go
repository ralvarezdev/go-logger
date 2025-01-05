package status

type (
	// Status is the status of the logger
	Status int
)

const (
	Info Status = iota
	Debug
	Warning
	Error
	Critical
)

// String returns the string representation of the status
func (s Status) String() string {
	switch s {
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	case Critical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}
