package go_logger

type (
	// Logger is an interface for logging messages
	Logger interface {
		Log(message *Message)
		Info(header, subheader string, details *[]string)
		Error(header, subheader string, errors *[]error)
		Debug(header, subheader string, details *[]string)
		Critical(header, subheader string, details *[]string)
		Warning(header, subheader string, details *[]string)
	}
)
