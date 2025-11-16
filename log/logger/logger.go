package logger

import (
	"context"
	"log/slog"
)

func NewLevelLogger(h slog.Handler) *LevelLogger {
	return &LevelLogger{
		Logger: slog.New(h),
	}
}

// LevelLogger extends slog.LevelLogger with additional log levels, Trace and Fatal.
type LevelLogger struct {
	*slog.Logger
}

// Trace logs a message at the Trace level, including optional attributes.
func (l *LevelLogger) Trace(msg string, attrs ...any) {
	l.TraceContext(context.Background(), msg, attrs...)
}

func (l *LevelLogger) TraceContext(ctx context.Context, msg string, attrs ...any) {
	l.Log(ctx, LevelTrace, msg, attrs...)
}

// Fatal logs a message at the Fatal level, including optional attributes.
func (l *LevelLogger) Fatal(msg string, attrs ...any) {
	l.FatalContext(context.Background(), msg, attrs...)
}

func (l *LevelLogger) FatalContext(ctx context.Context, msg string, attrs ...any) {
	l.Log(ctx, LevelFatal, msg, attrs...)
}
