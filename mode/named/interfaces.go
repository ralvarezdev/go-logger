package named

type (
	// Logger provides methods for logging messages at various severity levels for a specific mode.
	Logger interface {
		// Info logs an informational message.
		//
		// Parameters:
		//   subheader - a short description of the message.
		//   details - optional additional details.
		Info(subheader string, details ...string)
		// Error logs an error message.
		//
		// Parameters:
		//   subheader - a short description of the error.
		//   errors - the error values to log.
		Error(subheader string, errors ...error)
		// Debug logs a debug message.
		//
		// Parameters:
		//   subheader - a short description of the message.
		//   details - optional additional details.
		Debug(subheader string, details ...string)
		// Critical logs a critical message.
		//
		// Parameters:
		//   subheader - a short description of the message.
		//   details - optional additional details.
		Critical(subheader string, details ...string)
		// Warning logs a warning message.
		//
		// Parameters:
		//   subheader - a short description of the message.
		//   details - optional additional details.
		Warning(subheader string, details ...string)
	}
)
