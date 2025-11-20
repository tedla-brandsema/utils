package handler

import (
	"context"
	"log/slog"
	"runtime"
	"strings"
	"sync"

	"github.com/tedla-brandsema/utils/log/register"
)

type PkgAwareHandler struct {
	skip     int
	fallback slog.Handler
}

func NewPkgAwareHandler(fallback slog.Handler) *PkgAwareHandler {
	return &PkgAwareHandler{
		fallback: fallback,
	}
}

func (h *PkgAwareHandler) WithSkip(frames int) *PkgAwareHandler {
	return &PkgAwareHandler{
		fallback: h.fallback,
		skip:     frames,
	}
}

func (h *PkgAwareHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	// mode Pkg always returns ture
	if register.Mode() == register.Pkg {
		return true
	}

	// mode App
	return lvl >= register.ApplicationLevel()
}

func (h *PkgAwareHandler) Handle(ctx context.Context, r slog.Record) error {
	pcs := make([]uintptr, 1)
	runtime.Callers(4+h.skip, pcs)
	pc := pcs[0]
	r.PC = pc

	if lv, ok := pcLevelCache.Load(pc); ok {
		if r.Level >= lv.(*slog.LevelVar).Level() {
			return h.fallback.Handle(ctx, r)
		}
	}

	lv := resolveLevelVarForPC(pc, register.ApplicationLevel())
	pcLevelCache.Store(pc, lv)

	if r.Level >= lv.Level() {
		return h.fallback.Handle(ctx, r)
	}
	return nil
}

func (h *PkgAwareHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PkgAwareHandler{
		fallback: h.fallback.WithAttrs(attrs),
	}
}

func (h *PkgAwareHandler) WithGroup(name string) slog.Handler {
	return &PkgAwareHandler{
		fallback: h.fallback.WithGroup(name),
	}
}

var pcLevelCache sync.Map // pc uintptr → *slog.LevelVar

func resolveLevelVarForPC(pc uintptr, global slog.Level) *slog.LevelVar {
	// Find package (slow)
	pkg := packageFromPC(pc)

	// Already registered in Routes?
	if lv, ok := register.Routes.Get(pkg); ok {
		return lv
	}

	lv := register.AutoRegisterPackage(pkg, global)

	return lv
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
	// fast path: no slash → find first dot
	lastSlash := strings.LastIndexByte(full, '/')
	if lastSlash == -1 {
		if dot := strings.IndexByte(full, '.'); dot != -1 {
			return full[:dot]
		}
		return full
	}

	// after final slash, drop function name and method signature
	remainder := full[lastSlash+1:]

	// find first dot AFTER the slash
	dot := strings.IndexByte(remainder, '.')
	if dot != -1 {
		return full[:lastSlash+1+dot]
	}

	return full[:lastSlash]
}
