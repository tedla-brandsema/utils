package test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/tedla-brandsema/utils/log"
	"github.com/tedla-brandsema/utils/log/handler"
	"github.com/tedla-brandsema/utils/log/level"
)

func devHandlerOpts() (*slog.LevelVar, *handler.DevHandlerOptions) {
	opts := handler.NewDevHandlerOptions()
	opts.Color = true
	opts.AddSource = true
	lvlvar := opts.Level.(*slog.LevelVar) // Default is slog.LevelInfo (i.e. 0)
	return lvlvar, opts
}

func Test_DevHandler_Standard_Logger(t *testing.T) {
	lvlvar, opts := devHandlerOpts()
	devHandler := handler.NewDevHandler(os.Stdout, opts)

	lgr := slog.New(devHandler)
	slog.SetDefault(lgr)

	PrintLevels(lvlvar, slog.Default(), level.LogLevels, t)
}

func Test_DevHandler_Custom_Logger(t *testing.T) {
	lvlvar, opts := devHandlerOpts()
	devHandler := handler.NewDevHandler(os.Stdout, opts)

	lgr := log.NewLogger(devHandler)
	PrintLevels(lvlvar, lgr.Logger, level.LogLevels, t)
}

func Test_Set_Additional_Levels_JSONHandler(t *testing.T) {
	lvlvar := &slog.LevelVar{}
	lgr := slog.New(slog.NewJSONHandler(os.Stdout,
		level.SetAdditionalLogLevels(&slog.HandlerOptions{
			Level: lvlvar,
		}),
	))
	slog.SetDefault(lgr)

	PrintLevels(lvlvar, slog.Default(), level.LogLevels, t)
}

func Test_Set_Additional_Levels_TextHandler(t *testing.T) {
	lvlvar := &slog.LevelVar{}
	lgr := slog.New(slog.NewTextHandler(os.Stdout,
		level.SetAdditionalLogLevels(&slog.HandlerOptions{
			Level: lvlvar,
		}),
	))
	slog.SetDefault(lgr)

	PrintLevels(lvlvar, slog.Default(), level.LogLevels, t)
}

func PrintLevels(lvlvar *slog.LevelVar, lgr *slog.Logger, logLevels map[slog.Level]string, t *testing.T) {
	for lvl, name := range logLevels {
		lvlvar.Set(lvl)

		t.Logf("setting log level to: %s", name)

		lgr.Log(
			context.Background(),
			level.Trace,
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
			context.Background(),
			level.Fatal,
			"Unrecoverable error, shutting down",
			slog.String("service", "database connection"),
		)
	}
}
