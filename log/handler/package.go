package handler

import (
	"context"
	"log/slog"
	"runtime"
	"strings"
	"sync"

	"github.com/tedla-brandsema/utils/log"
)

type PkgAwareHandler struct {
    fallback slog.Handler
    minLevel slog.Level
}

func NewPkgAwareHandler(fallback slog.Handler, min slog.Level) *PkgAwareHandler {
    return &PkgAwareHandler{
        fallback: fallback,
        minLevel: min,
    }
}

func (h *PkgAwareHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return lvl >= h.minLevel
}

func (h *PkgAwareHandler) Handle(ctx context.Context, r slog.Record) error {
    pc := r.PC
    if pc == 0 {
        pcs := make([]uintptr, 1)
        runtime.Callers(5, pcs)
        pc = pcs[0]
    }


    pkg := packageFromPC(pc)
    if lvlv, ok := log.Routes.Get(pkg); ok {
        // Only emit if record level >= pkg-level
        if r.Level >= lvlv.Level() {
            return h.fallback.Handle(ctx, r)
        }
        return nil
    }

	// Auto-register new package with default level
	lvlv := &slog.LevelVar{}
	lvlv.Set(h.minLevel)
	log.Routes.Set(pkg, lvlv)

    return h.fallback.Handle(ctx, r)
}

func (h *PkgAwareHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PkgAwareHandler{
		fallback: h.fallback.WithAttrs(attrs),
		minLevel: h.minLevel,
	}
}

func (h *PkgAwareHandler) WithGroup(name string) slog.Handler {
	return &PkgAwareHandler{
		fallback: h.fallback.WithGroup(name),
		minLevel: h.minLevel,
	}
}

var pcCache sync.Map // pc → pkg string

func packageFromPC(pc uintptr) string {
	if v, ok := pcCache.Load(pc); ok {
		return v.(string)
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		pcCache.Store(pc, "")
		return ""
	}

	name := fn.Name()
	pkg := extractPkg(name)

	pcCache.Store(pc, pkg)
	return pkg
}

func extractPkg(full string) string {
	// fast manual scan instead of strings functions
	last := strings.LastIndexByte(full, '/')
	if last == -1 {
		last = 0
	}
	dot := strings.IndexByte(full[last:], '.')
	if dot == -1 {
		return full
	}
	return full[:last+dot]
}
