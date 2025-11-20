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
	fallback        slog.Handler
	globalThreshold *slog.LevelVar
	skip            int
}

func NewPkgAwareHandler(fallback slog.Handler, threshold *slog.LevelVar) *PkgAwareHandler {
	lvlv := threshold
	if lvlv == nil {
		lvlv = &slog.LevelVar{}
	}

	return &PkgAwareHandler{
		fallback:        fallback,
		globalThreshold: lvlv,
	}
}

func (h *PkgAwareHandler) WithSkip(frames int) *PkgAwareHandler {
	return &PkgAwareHandler{
		fallback:        h.fallback,
		globalThreshold: h.globalThreshold,
		skip:            frames,
	}
}

func (h *PkgAwareHandler) SetThreshold(lvl slog.Level) {
	h.globalThreshold.Set(lvl)
}

func (h *PkgAwareHandler) GetThreshold() slog.Level {
	return h.globalThreshold.Level()
}

func (h *PkgAwareHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return lvl >= h.globalThreshold.Level()
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

	lv := resolveLevelVarForPC(pc, h.globalThreshold)
	pcLevelCache.Store(pc, lv)

	return h.fallback.Handle(ctx, r)
}

func (h *PkgAwareHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PkgAwareHandler{
		fallback:        h.fallback.WithAttrs(attrs),
		globalThreshold: h.globalThreshold,
	}
}

func (h *PkgAwareHandler) WithGroup(name string) slog.Handler {
	return &PkgAwareHandler{
		fallback:        h.fallback.WithGroup(name),
		globalThreshold: h.globalThreshold,
	}
}

var pcLevelCache sync.Map // pc uintptr → *slog.LevelVar

func resolveLevelVarForPC(pc uintptr, global *slog.LevelVar) *slog.LevelVar {
	// Find package (slow)
	pkg := packageFromPC(pc)

	// Already registered in Routes?
	if lv, ok := register.Routes.Get(pkg); ok {
		return lv
	}

	lv := register.AutoRegisterPackage(pkg, global.Level())

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
