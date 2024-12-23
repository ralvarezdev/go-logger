package status

type (
	// Status is the status of the logger
	Status int
)

const (
	StatusSuccess Status = iota
	StatusFailed
	StatusError
	StatusWarning
	StatusInfo
	StatusDebug
	StatusTrace
	StatusNone
	StatusUnknown
)

// String returns the string representation of the status
func (s Status) String() string {
	switch s {
	case StatusSuccess:
		return "SUCCESS"
	case StatusFailed:
		return "FAILED"
	case StatusError:
		return "ERROR"
	case StatusWarning:
		return "WARNING"
	case StatusInfo:
		return "INFO"
	case StatusDebug:
		return "DEBUG"
	case StatusTrace:
		return "TRACE"
	case StatusNone:
		return ""
	default:
		return "UNKNOWN"
	}
}
