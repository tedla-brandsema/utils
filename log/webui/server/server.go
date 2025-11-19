package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tedla-brandsema/utils/log"
	"github.com/tedla-brandsema/utils/log/handler"
	"github.com/tedla-brandsema/utils/log/webui"
)

const port = ":8585"

func main() {
	opts := handler.NewDevHandlerOptions()
	opts.AddSource = true
	h := handler.NewDevHandler(os.Stdout, opts)

	// ph := handler.NewPkgAwareHandler(h, &slog.LevelVar{})
	// lgr := log.NewLogger(ph)
	// lgr := slog.New(ph)

	log.Set(os.Stdout, h)
	log.Info("mounting package level GUI")

	mux := http.NewServeMux()
	webui.Mount(mux)

	log.Info(fmt.Sprintf("serving on port %s", port))
	if err := http.ListenAndServe(port, mux); err != nil {
		panic(err)
	}
}
