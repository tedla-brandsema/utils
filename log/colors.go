package log

import "fmt"

// https://gist.github.com/JBlond/2fea43a3049b38287e5e9cefc87b2124

const ( //var clrStr = "\033[%d;%dm%s\033[0m"
	bgClrStr = "\033[%dm"
	clrStr   = "\033[%d;%dm"
	resetStr = "\033[0m"
	argsStr  = "%s"
)

type Format int

const (
	FormatRegular   Format = 0
	FormatBold      Format = 1
	FormatUnderline Format = 4
)

type Color int

// Base Colors
const (
	ColorDefault Color = 0

	ColorBlack Color = 30 + iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorPurple
	ColorCyan
	ColorWhite
)

// Intense Colors
const (
	IntenseColorBlack Color = 90 + iota
	IntenseColorRed
	IntenseColorGreen
	IntenseColorYellow
	IntenseColorBlue
	IntenseColorPurple
	IntenseColorCyan
	IntenseColorWhite
)

func styleStr(format Format, clr Color) string {
	return fmt.Sprintf(clrStr, format, clr)
}

func Style(format Format, color Color) func(...any) string {
	s := fmt.Sprint(styleStr(format, color), argsStr, resetStr)
	sprint := func(args ...any) string {
		return fmt.Sprintf(s, fmt.Sprint(args...))
	}
	return sprint
}

var (
	Red    = Style(FormatBold, IntenseColorRed)
	Green  = Style(FormatBold, IntenseColorGreen)
	Yellow = Style(FormatBold, IntenseColorYellow)
	Blue   = Style(FormatBold, IntenseColorBlue)
	Purple = Style(FormatBold, IntenseColorPurple)
	Cyan   = Style(FormatBold, IntenseColorCyan)
	White  = Style(FormatBold, IntenseColorWhite)
)
