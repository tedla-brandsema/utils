package log

import (
	"log/slog"
	"os"
	"testing"

)

func Test_Set_Additional_Levels_JSONHandler(t *testing.T) {
	lgr := slog.New(slog.NewJSONHandler(os.Stdout, SetAdditionalLogLevels(opts.HandlerOptions)))
	slog.SetDefault(lgr)

	PrintLevels(slog.Default(), t)
}

func Test_Set_Additional_Levels_TextHandler(t *testing.T) {
	lgr := slog.New(slog.NewTextHandler(os.Stdout, SetAdditionalLogLevels(opts.HandlerOptions)))
	slog.SetDefault(lgr)

	PrintLevels(slog.Default(), t)
}
