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
		return "Success"
	case StatusFailed:
		return "Failed"
	case StatusError:
		return "Error"
	case StatusWarning:
		return "Warning"
	case StatusInfo:
		return "Info"
	case StatusDebug:
		return "Debug"
	case StatusTrace:
		return "Trace"
	case StatusNone:
		return ""
	default:
		return "Unknown"
	}
}
