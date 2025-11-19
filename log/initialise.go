package log

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/tedla-brandsema/utils/log/handler"
	"github.com/tedla-brandsema/utils/log/level"
)

var (
	instance *Logger
	lvl      *slog.LevelVar
	out      io.Writer
	base     slog.Handler
)


func Set(w io.Writer, h slog.Handler) {
	lvl = &slog.LevelVar{}

	out = w
	if out == nil {
		out = os.Stdout
	}

	base = h
	if base == nil {
		base = slog.NewTextHandler(out, nil)
	}

	// ph := handler.NewPkgAwareHandler(base, lvl).WithSkip(1)
	ph := handler.NewPkgAwareHandler(base, lvl).WithSkip(1)
	instance = NewLogger(ph)
}


func GlobalThreshold(l slog.Level) {
	lvl.Set(l)
}

func Log(ctx context.Context, level slog.Level, msg string, args ...any) {
	instance.Log(ctx, level, msg, args...)
}

func Trace(msg string, attrs ...any) {
	instance.TraceContext(context.Background(), msg, attrs...)
}

func TraceContext(ctx context.Context, msg string, attrs ...any) {
	instance.Log(ctx, level.Trace, msg, attrs...)
}

func Debug(msg string, attrs ...any) {
	instance.DebugContext(context.Background(), msg, attrs...)
}

func DebugContext(ctx context.Context, msg string, attrs ...any) {
	instance.Log(ctx, level.Debug, msg, attrs...)
}

func Info(msg string, attrs ...any) {
	instance.InfoContext(context.Background(), msg, attrs...)
}

func InfoContext(ctx context.Context, msg string, attrs ...any) {
	instance.Log(ctx, level.Info, msg, attrs...)
}
func Warn(msg string, attrs ...any) {
	instance.WarnContext(context.Background(), msg, attrs...)
}

func WarnContext(ctx context.Context, msg string, attrs ...any) {
	instance.Log(ctx, level.Warn, msg, attrs...)
}
func Error(msg string, attrs ...any) {
	instance.ErrorContext(context.Background(), msg, attrs...)
}

func ErrorContext(ctx context.Context, msg string, attrs ...any) {
	instance.Log(ctx, level.Error, msg, attrs...)
}

func Fatal(msg string, attrs ...any) {
	instance.FatalContext(context.Background(), msg, attrs...)
}

func FatalContext(ctx context.Context, msg string, attrs ...any) {
	instance.Log(ctx, level.Fatal, msg, attrs...)
}
