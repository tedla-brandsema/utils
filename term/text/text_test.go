package text

import (
	"fmt"
	"github.com/tedla-brandsema/utils/term/color"
	"github.com/tedla-brandsema/utils/term/style"
	"testing"
)

func TestFormat(t *testing.T) {

	fmt.Println(
		Format("Hello World",
			style.Bold,
			style.Italic,
			color.Color(5),
		),
	)

}
