package log

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/tedla-brandsema/utils/generics"
)


var (
	packageLevels *generics.Registry[string, *slog.LevelVar] 
)

func Init() {
	packageLevels = generics.NewRegistry[string, *slog.LevelVar]() 
}

type PkgLvlHandler struct {
	Pkg string
	base slog.Handler
	opts *slog.HandlerOptions
}

func NewPkgLvlHandlerd(h slog.Handler, opts *slog.HandlerOptions) *PkgLvlHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	if h ==  nil {
		h = slog.NewTextHandler(os.Stdout, opts)
	}

	lvlVar := &slog.LevelVar{}
	if lvl := opts.Level; lvl != nil {
		lvlVar.Set(lvl.Level())

	}
	opts.Level = lvlVar

	pkg := callerPackage()
	packageLevels.Set(pkg, lvlVar)

	return &PkgLvlHandler{
		Pkg: pkg,
		base: h,
		opts: opts,
	}
}

func (h *PkgLvlHandler) Enabled(ctx context.Context, l slog.Level) bool {
     minLevel := slog.LevelInfo
    if h.opts.Level != nil {
        minLevel = h.opts.Level.Level()
    }
    return l >= minLevel    
}

func (h *PkgLvlHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.base.Handle(ctx, r) 
}

func (h *PkgLvlHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PkgLvlHandler{
		Pkg: h.Pkg,
		base: h.base.WithAttrs(attrs),
		opts: h.opts,
	}
}

func (h *PkgLvlHandler) WithGroup(name string) slog.Handler {
	return &PkgLvlHandler{
        Pkg:  h.Pkg,
        base: h.base.WithGroup(name),
        opts: h.opts,
    }
}


// callerPackage the returns the package of the caller.
// WARNING: is call stack dependant!
func callerPackage() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}
	name := fn.Name()
	parts := strings.Split(name, "/")
	last := parts[len(parts)-1]
	if idx := strings.Index(last, "."); idx != -1 {
		last = last[:idx]
	}
	return last
}

func SetPackageLevel(pkg string, lvl slog.Level) error {
	if lvlv, ok := packageLevels.Get(pkg); ok {
		lvlv.Set(lvl)
		return nil
	}
	return fmt.Errorf("unregistered package %q", pkg)
}

func Packages() map[string]slog.Level {
	pkgLvls := make(map[string]slog.Level)
	for k, v := range packageLevels.All() {
		pkgLvls[k] = v.Level()
	}
	return pkgLvls
}