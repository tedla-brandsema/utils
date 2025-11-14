package webui

import (
	"embed"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/tedla-brandsema/utils/log"
)
var (
	//go:embed templates/*.tmpl
	TemplatesFS embed.FS
	indexTmpl  = template.Must(template.ParseFS(TemplatesFS, "templates/index.html"))
	updateTmpl = template.Must(template.ParseFS(TemplatesFS, "templates/update.html"))
)

func listHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Packages []struct {
			Name  string
			Level string
		}
	}{}

	for pkg, lvl := range log.Packages()  {
			data.Packages = append(data.Packages, struct {
				Name  string
				Level string
			}{
				Name:  pkg,
				Level: lvl.Level().String(),
			})
	}

	indexTmpl.Execute(w, data)
}

func updateGetHandler(w http.ResponseWriter, r *http.Request) {
	pkg := r.URL.Query().Get("pkg")
	updateTmpl.Execute(w, struct{ Pkg string }{Pkg: pkg})
}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	pkg := r.FormValue("pkg")
	level := r.FormValue("level")
	// cascade := r.FormValue("cascade") == "1"

	var lvl slog.Level
	switch level {
	case "DEBUG":
		lvl = slog.LevelDebug
	case "INFO":
		lvl = slog.LevelInfo
	case "WARN":
		lvl = slog.LevelWarn
	default:
		lvl = slog.LevelError
	}

	// if cascade {
	// 	registry.SetCascadeLevel(pkg, lvl)
	// } else {
		log.SetPackageLevel(pkg, lvl)
	// }

	http.Redirect(w, r, "/logdev", http.StatusSeeOther)
}

// Mount attaches all handlers to a mux.
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/log/level", listHandler)
	mux.HandleFunc("/log/level/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			updatePostHandler(w, r)
			return
		}
		updateGetHandler(w, r)
	})
}
