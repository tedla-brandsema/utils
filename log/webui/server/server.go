package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/tedla-brandsema/utils/log/handler"
	"github.com/tedla-brandsema/utils/log/webui"
)

const port =":8585"

func main() {
    opts := slog.HandlerOptions{
    	Level:     &slog.LevelVar{},
    }
    // h := handler.NewDevHandler(os.Stdout, handler.NewDevHandlerOptions())
    ph := handler.NewPkgLvlHandlerd(slog.NewTextHandler(os.Stdout, &opts), &opts)
    lgr := slog.New(ph)

    // Register your own packages:
    // handler.Register("myapp")
    // handler.Register("myapp.db")
    // handler.Register("myapp.http")
    lgr.Info("mounting package level GUI")
    
    mux := http.NewServeMux()
    webui.Mount(mux)

    lgr.Info(fmt.Sprintf("serving on port %s", port))
    if err := http.ListenAndServe(port, mux); err != nil {
        panic(err)
    }
}