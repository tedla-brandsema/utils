package log

import (
	"context"
	"log/slog"
	"testing"
)

var (
	lvl  *slog.LevelVar
	opts *DevHandlerOptions
	ctx  = context.Background()
)

func init() {
	opts = NewDevHandlerOptions()
	opts.Color = true
	opts.AddSource = true
	lvl = opts.Level.(*slog.LevelVar) // Default is slog.LevelInfo (i.e. 0)
}

var levels = []slog.Level{
	LevelTrace,
	LevelDebug,
	LevelInfo,
	LevelWarn,
	LevelError,
	LevelFatal,
}

func PrintLevels(lgr *slog.Logger, t *testing.T) {
	for _, level := range levels {
		lvl.Set(level)

		t.Logf("setting log level to: %s", LevelString(level))

		lgr.Log(
			ctx,
			LevelTrace,
			"Starting function execution",
			slog.String("function_name", "processData"),
		)
		lgr.Debug(
			"executing database query",
			slog.String("query", "SELECT * FROM users"),
		)
		lgr.Info(
			"image upload successful",
			slog.String("image_id", "39ud88"),
		)
		lgr.Warn(
			"storage is 90% full",
			slog.String("available_space", "900.1 MB"),
		)
		lgr.Error(
			"An error occurred while processing the request",
			slog.String("url", "https://example.com"),
		)
		lgr.Log(
			ctx,
			LevelFatal,
			"Unrecoverable error, shutting down",
			slog.String("service", "database connection"),
		)
	}
}
