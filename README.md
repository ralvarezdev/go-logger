# go-logger

Logger for Go projects with severity statuses, a mode-aware logger (output depends on the application mode flag) and a "named" logger that prefixes entries with a fixed header. Requires Go 1.23.4 (per `go.mod`).

**Note:** This repository is archived and read-only.

## Installation

```bash
go get github.com/ralvarezdev/go-logger
```

Direct dependencies: `go-flags` and `go-strings` from `github.com/ralvarezdev`.

## Packages

- **`go_logger`** (root) — `Logger` interface (`Log(*Message)`, `Info`, `Error`, `Debug`, `Critical`, `Warning`, each taking `header, subheader` plus details or errors), `NewMessage(...)`, `NewDefaultLogger()`.
- **`status`** — `Status` enum: `Info`, `Debug`, `Warning`, `Error`, `Critical`.
- **`mode`** — `Logger` interface adding `ShouldLog(status)` and `RunIfShouldLog(status, fn)`; `NewDefaultLogger(logger, flagMode)` wraps any `go_logger.Logger` and filters by the `go-flags` mode flag. The per-mode rules are in `mode/logger.go`.
- **`mode/named`** — `Logger` without header (`Info(subheader, details...)`, ...); `NewDefaultLogger(header, logger)` wraps a `mode.Logger` with a fixed header.

## Usage (sketch)

```go
modeFlag := goflagsmode.NewFlag(goflagsmode.DefaultMode, goflagsmode.AllowedModes)

base := gologger.NewDefaultLogger()
modeLogger, err := gologgermode.NewDefaultLogger(base, modeFlag)
named, err := gologgermodenamed.NewDefaultLogger("MyService", modeLogger)

named.Info("started", "listening on :8080")
```

## Development

```bash
go build ./...
```

There are no tests.

## License

GNU General Public License v3.0 (see [LICENSE](LICENSE)).
