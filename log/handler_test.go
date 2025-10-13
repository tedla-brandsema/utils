package log

import (
	"log/slog"
	"testing"
)

func Test_DevHandler_Standard_Logger(t *testing.T) {
	lgr := slog.New(NewDevHandler(out, opts))
	slog.SetDefault(lgr)

	printLevels(slog.Default(), t)
}

func Test_DevHandler_Custom_Logger(t *testing.T) {
	lgr := New(out, opts)
	printLevels(lgr.Logger, t)
}
