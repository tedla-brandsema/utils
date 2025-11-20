package main

import (
	"io"
	"log/slog"
	"os"

	"github.com/tedla-brandsema/utils/log/handler"
	"github.com/tedla-brandsema/utils/log/level"
	"github.com/tedla-brandsema/utils/log/webui"
)



func slogSetup(w io.Writer) {
	opts := handler.NewDevHandlerOptions()
	opts.LevelVar().Set(level.Trace)
	dh := handler.NewDevHandler(w, opts)
	ph := handler.NewPkgAwareHandler(dh, opts.LevelVar()).WithSkip(1)

	lgr := slog.New(ph)
	slog.SetDefault(lgr)
}

func main() {
	slogSetup(os.Stdout)
	webui.Start(":8585") 
}
