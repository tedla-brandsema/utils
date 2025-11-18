package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/tedla-brandsema/utils/log/handler"
	"github.com/tedla-brandsema/utils/log/logger"
	"github.com/tedla-brandsema/utils/log/webui"
)

const port =":8585"

func main() {
    
    h := handler.NewDevHandler(os.Stdout, handler.NewDevHandlerOptions())
    ph := handler.NewPkgAwareHandler(h, logger.LevelTrace )
    lgr := slog.New(ph)

    lgr.Info("mounting package level GUI")
    
    mux := http.NewServeMux()
    webui.Mount(mux)

    lgr.Info(fmt.Sprintf("serving on port %s", port))
    if err := http.ListenAndServe(port, mux); err != nil {
        panic(err)
    }
}