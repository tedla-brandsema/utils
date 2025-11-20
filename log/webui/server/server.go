package main

import (
	"context"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/tedla-brandsema/utils/log/handler"
	"github.com/tedla-brandsema/utils/log/level"
	"github.com/tedla-brandsema/utils/log/webui"
)

func slogSetup(w io.Writer) {
	opts := handler.NewDevHandlerOptions()
	opts.LevelVar().Set(level.Trace)
	dh := handler.NewDevHandler(w, opts)
	ph := handler.NewPkgAwareHandler(dh, opts.LevelVar()).WithSkip(0)

	lgr := slog.New(ph)
	slog.SetDefault(lgr)
}

func main() {
	slogSetup(os.Stdout)
	shutdown, err := webui.Start(":8585")
	if err != nil {
		panic(err)
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	<-ctx.Done()
	shutdown(context.Background())
}
