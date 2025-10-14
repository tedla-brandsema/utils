package log

import (
	"fmt"
	"log/slog"
)

// Custom log levels for additional verbosity control.
const (
	// NOTE: do not set to slog.LevelWarn directly;
	// always calc the delta
	step = slog.LevelWarn - slog.LevelInfo

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

var defaultPaddedLevelText = map[slog.Level]string{
	LevelTrace: padLevelText(LevelTrace),
	LevelDebug: padLevelText(LevelDebug),
	LevelInfo:  padLevelText(LevelInfo),
	LevelWarn:  padLevelText(LevelWarn),
	LevelError: padLevelText(LevelError),
	LevelFatal: padLevelText(LevelFatal),
}

var defaultColoredLevelText = map[slog.Level]string{
	LevelTrace: Blue(defaultPaddedLevelText[LevelTrace]),
	LevelDebug: Green(defaultPaddedLevelText[LevelDebug]),
	LevelInfo:  White(defaultPaddedLevelText[LevelInfo]),
	LevelWarn:  Yellow(defaultPaddedLevelText[LevelWarn]),
	LevelError: Red(defaultPaddedLevelText[LevelError]),
	LevelFatal: Purple(defaultPaddedLevelText[LevelFatal]),
}

var paddingStr = fmt.Sprintf("%%-%ds", calcPadding())

func padLevelText(lvl slog.Level) string {
	return fmt.Sprintf(paddingStr, defaultLevelText[lvl])
}

func calcPadding() int {
	width := 0
	for _, v := range defaultLevelText {
		if len(v) > width {
			width = len(v)
		}
	}
	return width
}

func LevelString(level slog.Level) string {
	return defaultPaddedLevelText[level]
}

func ColoredLevelString(level slog.Level) string {
	return defaultColoredLevelText[level]
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

	if opts.Level != nil {
		lvl.Set(opts.Level.Level())
	}

	opts.Level = lvl
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
