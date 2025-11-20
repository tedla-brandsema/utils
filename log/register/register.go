package register

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/tedla-brandsema/utils/generics"
)

type LevelMode int

const (
	App LevelMode = iota
	Pkg
)

var (
	lvlMode = Pkg
	global  = &slog.LevelVar{}

	modeLock sync.RWMutex
	lvlLock  sync.RWMutex
)

func Mode() LevelMode {
	modeLock.RLock()
	defer modeLock.RUnlock()

	return lvlMode
}

func SetMode(mode LevelMode) {
	modeLock.Lock()
	defer modeLock.Unlock()

	lvlMode = mode
}

func ApplicationLevel() slog.Level {
	lvlLock.RLock()
	defer lvlLock.RUnlock()

	return global.Level()
}

func SetApplicationLevel(lvl slog.Level) {
	lvlLock.Lock()
	defer lvlLock.Unlock()

	global.Set(lvl)
}

// JSON persistence config
var (
	persistFile string
	persistMu   sync.Mutex
)

func EnableJSONPersistence(filePath string) error {
	persistFile = filePath
	return loadFromJSON()
}

func loadFromJSON() error {
	if persistFile == "" {
		return nil
	}
	data, err := os.ReadFile(persistFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var pkgMap map[string]int
	if err := json.Unmarshal(data, &pkgMap); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	for pkg, lvl := range pkgMap {
		lvlVar := &slog.LevelVar{}
		lvlVar.Set(slog.Level(lvl))
		Routes.Set(pkg, lvlVar)
	}
	return nil
}

func saveToJSON() error {
	if persistFile == "" {
		return nil
	}
	persistMu.Lock()
	defer persistMu.Unlock()

	all := Routes.All()
	pkgMap := make(map[string]int, len(all))
	for pkg, lvlv := range all {
		pkgMap[pkg] = int(lvlv.Level())
	}

	tmpFile := persistFile + ".tmp"
	data, _ := json.MarshalIndent(pkgMap, "", "  ")
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		return err
	}

	return os.Rename(tmpFile, persistFile)
}

func AutoRegisterPackage(pkg string, defaultLvl slog.Level) *slog.LevelVar {
	lvlv := &slog.LevelVar{}
	lvlv.Set(defaultLvl)
	Routes.Set(pkg, lvlv)
	if persistFile != "" {
		_ = saveToJSON()
	}
	return lvlv
}

var (
	Routes = generics.NewRegistry[string, *slog.LevelVar]()
)

func SetPackageLevel(pkg string, lvl slog.Level) error {
	if lvlvar, ok := Routes.Get(pkg); ok {
		lvlvar.Set(lvl)
		if persistFile != "" {
			return saveToJSON()
		}
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
