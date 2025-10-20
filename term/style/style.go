package style

import (
	"github.com/tedla-brandsema/utils/term/ansi"
	"strconv"
)

var styleIndex = map[int]struct{}{
	1: {},
	2: {},
	3: {},
	4: {},
	9: {},
}

func ValidStyle(s int) bool {
	_, ok := styleIndex[s]
	return ok
}

func Param(s int) ansi.Param {
	return ansi.Param(strconv.Itoa(s))
}

var (
	Bold      = Param(1)
	Faint     = Param(2)
	Italic    = Param(3)
	Underline = Param(4)
	Strike    = Param(9)
)
