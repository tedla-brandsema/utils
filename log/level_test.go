package log

import (
	"log/slog"
	"os"
	"testing"
)

func Test_Set_Additional_Levels_JSONHandler(t *testing.T) {
	lgr := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       lvl,
		ReplaceAttr: ReplaceAttrLevel,
	}))
	slog.SetDefault(lgr)

	printLevels(slog.Default(), t)
}

func Test_Set_Additional_Levels_TextHandler(t *testing.T) {
	lgr := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:       lvl,
		ReplaceAttr: ReplaceAttrLevel,
	}))
	slog.SetDefault(lgr)

	printLevels(slog.Default(), t)
}
