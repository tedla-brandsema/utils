package color

import (
	"fmt"
	"strconv"
)

func HexToRGBA(hex string) (r, g, b, a uint8, err error) {
	if hex[0] == '#' {
		hex = hex[1:]
	}
	if len(hex) != 6 && len(hex) != 8 {
		return 0, 0, 0, 255, fmt.Errorf("invalid hex length")
	}
	r64, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g64, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b64, _ := strconv.ParseUint(hex[4:6], 16, 8)
	a64 := uint64(255)
	if len(hex) == 8 {
		a64, _ = strconv.ParseUint(hex[6:8], 16, 8)
	}
	return uint8(r64), uint8(g64), uint8(b64), uint8(a64), nil
}

func ColorCubeRGB(code int) (r, g, b int) {
	if code < 16 || code > 231 {
		return 0, 0, 0
	}
	c := code - 16
	r = (c / 36) % 6
	g = (c / 6) % 6
	b = c % 6
	r = []int{0, 95, 135, 175, 215, 255}[r]
	g = []int{0, 95, 135, 175, 215, 255}[g]
	b = []int{0, 95, 135, 175, 215, 255}[b]
	return
}
