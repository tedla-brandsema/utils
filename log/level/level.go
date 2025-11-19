package level

import (
	"fmt"
	"log/slog"

	"github.com/tedla-brandsema/utils/log/color"
)

// Custom log levels for additional verbosity control.
const (
	// NOTE: do not set to slog.LevelWarn directly;
	// always calc the delta
	step = slog.LevelWarn - slog.LevelInfo

	Trace = slog.LevelDebug - step
	Debug = slog.LevelDebug
	Info  = slog.LevelInfo
	Warn  = slog.LevelWarn
	Error = slog.LevelError
	Fatal = slog.LevelError + step
)

var LogLevels = map[slog.Level]string{
	Trace: "TRACE",
	Debug: slog.LevelDebug.String(),
	Info:  slog.LevelInfo.String(),
	Warn:  slog.LevelWarn.String(),
	Error: slog.LevelError.String(),
	Fatal: "FATAL",
}

// FIXME: defaultPaddedLevelText, defaultColoredLevelText, padLevelText, calcPadding
// all rely on defaultLevelText, it should work dynamically based on LogLevels
var defaultPaddedLevelText = map[slog.Level]string{
	Trace: padLevelText(Trace),
	Debug: padLevelText(Debug),
	Info:  padLevelText(Info),
	Warn:  padLevelText(Warn),
	Error: padLevelText(Error),
	Fatal: padLevelText(Fatal),
}

var defaultColoredLevelText = map[slog.Level]string{
	Trace: color.Blue(defaultPaddedLevelText[Trace]),
	Debug: color.Green(defaultPaddedLevelText[Debug]),
	Info:  color.White(defaultPaddedLevelText[Info]),
	Warn:  color.Yellow(defaultPaddedLevelText[Warn]),
	Error: color.Red(defaultPaddedLevelText[Error]),
	Fatal: color.Purple(defaultPaddedLevelText[Fatal]),
}

var paddingStr = fmt.Sprintf("%%-%ds", calcPadding())

func padLevelText(lvl slog.Level) string {
	return fmt.Sprintf(paddingStr, LogLevels[lvl])
}

func calcPadding() int {
	width := 0
	for _, v := range LogLevels {
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
	if opts == nil {
		opts = &slog.HandlerOptions{
			Level: &slog.LevelVar{},
		}
	}
	opts.ReplaceAttr = replaceAttrLevel

	return opts
}

func replaceAttrLevel(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.LevelKey {
		level := a.Value.Any().(slog.Level)
		a.Value = slog.StringValue(LogLevels[level])
	}
	return a
}
