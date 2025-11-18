package log

import (
	"fmt"
	"log/slog"

	"github.com/tedla-brandsema/utils/generics"
)

var (
	Routes = generics.NewRegistry[string, *slog.LevelVar]()
)

func SetPackageLevel(pkg string, lvl slog.Level) error {
	if lvlvar, ok := Routes.Get(pkg); ok {
		lvlvar.Set(lvl)
		return nil
	}
	return fmt.Errorf("unregistered package %q", pkg)
}

func Packages() map[string]slog.Level {
	pkgLvls := make(map[string]slog.Level)
	for k, v := range Routes.All() {
		pkgLvls[k] = v.Level()
	}
	return pkgLvls
}
