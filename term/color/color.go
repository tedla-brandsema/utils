package color

import (
	"fmt"
	"github.com/tedla-brandsema/utils/term/ansi"
)

// Select Graphic Rendition Colors

const (
	foregroundColorParam = "38;5;%d"
	backgroundColorParam = "48;5;%d"
)

func Color(c uint8) ansi.Param {
	return ansi.Param(fmt.Sprintf(foregroundColorParam, c))
}

func BackgroundColor(c uint8) ansi.Param {
	return ansi.Param(fmt.Sprintf(backgroundColorParam, c))
}

const (
	rgbForegroundColorParam = "38;2;%d;%d;%d"
	rgbBackgroundColorParam = "48;2;%d;%d;%d"
)

func RgbColor(r uint8, g uint8, b uint8) ansi.Param {
	return ansi.Param(fmt.Sprintf(rgbForegroundColorParam, r, g, b))
}

func RgbBackgroundColor(r uint8, g uint8, b uint8) ansi.Param {
	return ansi.Param(fmt.Sprintf(rgbBackgroundColorParam, r, g, b))
}

func ContrastColor(n uint8) uint8 {
	if n == 0 {
		return 15 // bright white on black
	}
	if n >= 232 { // grayscale ramp
		if n-232 < 12 {
			return 15
		}
		return 0
	}
	// RGB cube: rough luminance
	r := ((n - 16) / 36) % 6 * 51
	g := ((n - 16) / 6) % 6 * 51
	b := (n - 16) % 6 * 51
	luma := 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)
	if luma < 128 {
		return 15
	}
	return 0
}

func Pallet() {
	for i := 0; i < 256; i++ {
		bg := BackgroundColor(uint8(i))
		if i%8 == 0 {
			fmt.Println()
		}
		contrast := Color(ContrastColor(uint8(i)))
		txt := ansi.SGRSequence(bg, contrast)(fmt.Sprintf(" %03d ", i))
		fmt.Print(txt)

		continue
	}
	fmt.Println()
}
