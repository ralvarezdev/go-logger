package status

type (
	// Status is the status of the logger
	Status int
)

const (
	Success Status = iota
	Failed
	Error
	Warning
	Info
	Debug
	Trace
	None
	Unknown
)

// String returns the string representation of the status
func (s Status) String() string {
	switch s {
	case Success:
		return "SUCCESS"
	case Failed:
		return "FAILED"
	case Error:
		return "ERROR"
	case Warning:
		return "WARNING"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	case Trace:
		return "TRACE"
	case None:
		return ""
	default:
		return "UNKNOWN"
	}
}
