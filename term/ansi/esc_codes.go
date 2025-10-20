package ansi

import (
	"strings"
)

// https://en.wikipedia.org/wiki/ANSI_escape_code

// ANSI escape codes
const (
	// Control Codes
	esc = "\033" // Escape

	// Escape Sequences
	csi = esc + "["  // Control Sequence Introducer
	st  = esc + "\\" // String Terminator
	osc = esc + "]"  // Operating System Command

	// CSI Sequences
	//sgr = csi + "%dm" // Select Graphic Rendition
)

// Control Sequence Introducer
// Parameters
const (
	sgrFinalByte = "m"

	resetParam = "0"

	Reset = csi + resetParam + sgrFinalByte
)

type Param string

type ParamFunc func(int) Param

func ParamMust(p Param, err error) Param {
	if err != nil {
		panic(err)
	}
	return p
}

func SGRSequence(param ...Param) func(string) string {
	sgr := csi + joinStrings(param, ";") + sgrFinalByte
	//sgr := csi + strings.Join(param, ";") + sgrFinalByte
	return func(s string) string {
		return sgr + s + Reset
	}
}

func joinStrings[T ~string](elems []T, sep string) string {
	if len(elems) == 0 {
		return ""
	}

	n := len(sep) * (len(elems) - 1)
	for _, e := range elems {
		n += len(e)
	}

	var b strings.Builder
	b.Grow(n)

	b.WriteString(string(elems[0]))
	for _, e := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(string(e))
	}
	return b.String()
}
