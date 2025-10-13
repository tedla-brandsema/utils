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

// setAdditionalLogLevels modifies slog.HandlerOptions to support custom log level names.
func setAdditionalLogLevels(opts *slog.HandlerOptions) {
	opts.ReplaceAttr = ReplaceAttrLevel
}

func ReplaceAttrLevel(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level := a.Value.Any().(slog.Level)
		a.Value = slog.StringValue(defaultLevelText[level])
	}
	return a
}
