package text

import "github.com/tedla-brandsema/utils/term/ansi"

func Format(txt string, param ...ansi.Param) string {
	return ansi.SGRSequence(param...)(txt)
}
