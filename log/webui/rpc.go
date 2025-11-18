package webui

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/tedla-brandsema/utils/log"
	"github.com/tedla-brandsema/utils/log/logger"
)





var levels []byte

func loadLevels() {
	var err error

	levelMap := make(map[int]string)
	for k, v := range logger.LogLevels {
		levelMap[int(k)] = v
	}

	levels, err = json.MarshalIndent(levelMap, "", "\t")
	if err != nil {
		panic(err)
	}
}

func levelHandler(w http.ResponseWriter, r *http.Request) {
	if len(levels) == 0 {
		loadLevels()
	}
	if _, err := w.Write(levels); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Package struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

func packageHandler(w http.ResponseWriter, r *http.Request) {
	var data []Package

	for pkg, lvl := range log.Packages() {
		data = append(data,
			Package{
				Name:  pkg,
				Level: int(lvl.Level()),
			})
	}

	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	pkg := r.FormValue("pkg")
	// cascade := r.FormValue("cascade") == "1"

	level, err := strconv.Atoi(r.FormValue("level"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var lvl slog.Level = slog.Level(level)

	if _, ok := logger.LogLevels[lvl]; !ok { // check validity of level
		http.Error(w, fmt.Errorf("invalid level index %d", lvl).Error(), http.StatusBadRequest)
		return
	}

	// if cascade {
	// 	registry.SetCascadeLevel(pkg, lvl)
	// } else {
	log.SetPackageLevel(pkg, lvl)
	// }

	// http.Redirect(w, r, basePath, http.StatusSeeOther)
}


// func globalThresholdHandler(w http.ResponseWriter, r *http.Request) {

// }

// func updateGlobalThresholdHandler(w http.ResponseWriter, r *http.Request) {

// }