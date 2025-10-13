package log

import (
	"fmt"
	"log/slog"
)

// Custom log levels for additional verbosity control.
const (
	step = 4

	LevelTrace = slog.LevelDebug - step
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
	LevelFatal = slog.LevelError + step
)

var defaultLevelText = map[slog.Level]string{
	LevelTrace: "TRACE",
	LevelDebug: slog.LevelDebug.String(),
	LevelInfo:  slog.LevelInfo.String(),
	LevelWarn:  slog.LevelWarn.String(),
	LevelError: slog.LevelError.String(),
	LevelFatal: "FATAL",
}

var defaultColoredLevelText = map[slog.Level]string{
	LevelTrace: Blue(defaultPaddedLevelText(LevelTrace)),
	LevelDebug: Green(defaultPaddedLevelText(LevelDebug)),
	LevelInfo:  White(defaultPaddedLevelText(LevelInfo)),
	LevelWarn:  Yellow(defaultPaddedLevelText(LevelWarn)),
	LevelError: Red(defaultPaddedLevelText(LevelError)),
	LevelFatal: Purple(defaultPaddedLevelText(LevelFatal)),
}

func defaultPaddedLevelText(lvl slog.Level) string {
	return fmt.Sprintf("%-5s", defaultLevelText[lvl])
}

var levelText map[slog.Level]string

func init() {
	levelText = defaultColoredLevelText
}

func LevelString(level slog.Level) string {
	return levelText[level]
}

// SetAdditionalLogLevels modifies slog.HandlerOptions to support custom log level names.
// It sets the current log level to slog.LevelInfo if non is set.
func SetAdditionalLogLevels(opts *slog.HandlerOptions) *slog.HandlerOptions {
	// To dynamically set the log level,
	// opts.Level needs to be set to a slog.LevelVar
	lvl := &slog.LevelVar{}

	if opts == nil {
		opts = &slog.HandlerOptions{
			Level: lvl,
		}
	}

	if opts.Level == nil {
		lvl.Set(opts.Level.Level())
		opts.Level = lvl
	}

	opts.ReplaceAttr = replaceAttrLevel

	return opts
}

func replaceAttrLevel(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level := a.Value.Any().(slog.Level)
		a.Value = slog.StringValue(defaultLevelText[level])
	}
	return a
}
