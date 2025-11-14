package log

import (
	"log/slog"
	"os"
	"testing"
)

var (
	out = os.Stdout
)

func Test_DevHandler_Standard_Logger(t *testing.T) {
	lgr := slog.New(NewDevHandler(out, opts))
	slog.SetDefault(lgr)

	PrintLevels(slog.Default(), t)
}

func Test_DevHandler_Custom_Logger(t *testing.T) {
	lgr := New(out, opts)
	PrintLevels(lgr.Logger, t)
}
