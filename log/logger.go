package log

import (
	"context"
	"io"
	"log/slog"
	"os"
)

// New initializes a new slog.Logger
func New(out io.Writer, opts *slog.HandlerOptions) *Logger {
	return &Logger{
		Logger: slog.New(NewDevHandler(out, opts)),
	}
}

func NewDefault() *Logger {
	return New(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     LevelInfo,
	})
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
	l.Log(ctx, LevelTrace, msg, attrs...)
}

// Fatal logs a message at the Fatal level, including optional attributes.
func (l *Logger) Fatal(msg string, attrs ...any) {
	l.FatalContext(context.Background(), msg, attrs...)
}

func (l *Logger) FatalContext(ctx context.Context, msg string, attrs ...any) {
	l.Log(ctx, LevelFatal, msg, attrs...)
}
