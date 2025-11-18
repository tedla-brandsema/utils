package handler

import (
	"context"
	"log/slog"
	"runtime"
	"strings"
	"sync"
)

type pkgRoute struct {
	handler slog.Handler
	level   slog.Level
}

type PkgAwareHandler struct {
	fallback slog.Handler

	// global minimum level
	minLevel slog.Level

	// map of pkg → handler + level
	mu     sync.RWMutex
	routes map[string]pkgRoute
}

func NewPkgAwareHandler(fallback slog.Handler, min slog.Level) *PkgAwareHandler {
	return &PkgAwareHandler{
		fallback: fallback,
		minLevel: min,
		routes:   make(map[string]pkgRoute),
	}
}

func (h *PkgAwareHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return lvl >= h.minLevel
}

// Register or update package-level handler
func (h *PkgAwareHandler) SetPackageHandler(pkg string, handler slog.Handler, level slog.Level) {
	h.mu.Lock()
	h.routes[pkg] = pkgRoute{handler: handler, level: level}
	h.mu.Unlock()
}

func (h *PkgAwareHandler) Handle(ctx context.Context, r slog.Record) error {
	// only do the expensive lookup here
	pc := r.PC
	if pc == 0 {
		pcs := make([]uintptr, 1)
		runtime.Callers(5, pcs)
		pc = pcs[0]
	}

	pkg := packageFromPC(pc)

	h.mu.RLock()
	route, ok := h.routes[pkg]
	h.mu.RUnlock()

	if ok && r.Level >= route.level {
		return route.handler.Handle(ctx, r)
	}

	// fallback handler
	return h.fallback.Handle(ctx, r)
}

func (h *PkgAwareHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PkgAwareHandler{
		fallback: h.fallback.WithAttrs(attrs),
		minLevel: h.minLevel,
		routes:   h.routes, // shared intentionally
	}
}

func (h *PkgAwareHandler) WithGroup(name string) slog.Handler {
	return &PkgAwareHandler{
		fallback: h.fallback.WithGroup(name),
		minLevel: h.minLevel,
		routes:   h.routes,
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
