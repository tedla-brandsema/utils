package log

import (
	"context"
	"log/slog"

	"github.com/tedla-brandsema/utils/log/level"
)

func NewLogger(h slog.Handler) *Logger {
	return &Logger{
		Logger: slog.New(h),
	}
}

// Logger extends slog.Logger with additional log levels, Trace and Fatal.
type Logger struct {
	*slog.Logger
}

// Trace logs a message at the Trace level, including optional attributes.
func (l *Logger) Trace(msg string, attrs ...any) {
	l.TraceContext(context.Background(), msg, attrs...)
}

func (l *Logger) TraceContext(ctx context.Context, msg string, attrs ...any) {
	l.Log(ctx, level.Trace, msg, attrs...)
}

// Fatal logs a message at the Fatal level, including optional attributes.
func (l *Logger) Fatal(msg string, attrs ...any) {
	l.FatalContext(context.Background(), msg, attrs...)
}

func (l *Logger) FatalContext(ctx context.Context, msg string, attrs ...any) {
	l.Log(ctx, level.Fatal, msg, attrs...)
}
