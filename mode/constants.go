package mode

import (
	goflagsmode "github.com/ralvarezdev/go-flags/mode"
	gologgerstatus "github.com/ralvarezdev/go-logger/status"
)

var (
	// LogModeMap is the map of the log mode
	LogModeMap = map[goflagsmode.Mode]map[gologgerstatus.Status]bool{
		goflagsmode.Debug: {
			gologgerstatus.Info:     true,
			gologgerstatus.Error:    true,
			gologgerstatus.Warning:  true,
			gologgerstatus.Debug:    true,
			gologgerstatus.Critical: true,
		},
		goflagsmode.Dev: {
			gologgerstatus.Info:     true,
			gologgerstatus.Error:    true,
			gologgerstatus.Warning:  true,
			gologgerstatus.Debug:    false,
			gologgerstatus.Critical: true,
		},
		goflagsmode.Prod: {
			gologgerstatus.Info:     true,
			gologgerstatus.Error:    true,
			gologgerstatus.Warning:  true,
			gologgerstatus.Debug:    false,
			gologgerstatus.Critical: true,
		},
	}
)
