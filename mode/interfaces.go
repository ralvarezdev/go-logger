package mode

import (
	gologger "github.com/ralvarezdev/go-logger"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
)

type (
	// Logger interface for the mode logger
	Logger interface {
		ShouldLog(status gologgerstatus.Status) bool
		RunIfShouldLog(status gologgerstatus.Status, fn func())
		gologger.Logger
	}
)
