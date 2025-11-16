package color

import "fmt"

// https://gist.github.com/JBlond/2fea43a3049b38287e5e9cefc87b2124

const ( //var clrStr = "\033[%d;%dm%s\033[0m"
	bgClrStr = "\033[%dm"
	clrStr   = "\033[%d;%dm"
	resetStr = "\033[0m"
	argsStr  = "%s"
)

type Style int

const (
	StyleRegular   Style = 0
	StyleBold      Style = 1
	StyleFaint     Style = 2
	StyleItalic    Style = 3
	StyleUnderline Style = 4
	StyleStrike    Style = 9
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

func styleStr(format Style, clr Color) string {
	return fmt.Sprintf(clrStr, format, clr)
}

func Format(style Style, color Color) func(...any) string {
	s := fmt.Sprint(styleStr(style, color), argsStr, resetStr)
	sprint := func(args ...any) string {
		return fmt.Sprintf(s, fmt.Sprint(args...))
	}
	return sprint
}

var (
	Red    = Format(StyleBold, IntenseColorRed)
	Green  = Format(StyleBold, IntenseColorGreen)
	Yellow = Format(StyleBold, IntenseColorYellow)
	Blue   = Format(StyleBold, IntenseColorBlue)
	Purple = Format(StyleBold, IntenseColorPurple)
	Cyan   = Format(StyleBold, IntenseColorCyan)
	White  = Format(StyleBold, IntenseColorWhite)
)
