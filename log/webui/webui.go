package webui

import (
	"embed"
	"html/template"
	"net/http"
	"net/url"
)

const (
	basePath = "/log/level"
)

var (
	//go:embed templates/*.tmpl
	TemplatesFS embed.FS

	indexTmpl  = template.Must(template.ParseFS(TemplatesFS, "templates/index.tmpl"))
	// updateTmpl = template.Must(template.ParseFS(TemplatesFS, "templates/update.tmpl"))
)


// Mount attaches all handlers to a mux.
func Mount(mux *http.ServeMux) {
	// Templates
	mux.HandleFunc("GET " + basePath, indexHandler)


	// RPC
	levelPath, _ := url.JoinPath(basePath, "all")
	mux.HandleFunc("GET " + levelPath, levelHandler)

	levelPackagePath, _ := url.JoinPath(basePath, "package", "all")
	mux.HandleFunc("GET " + levelPackagePath, packageHandler)

	updatePath, _ := url.JoinPath(basePath, "package", "update")
	mux.HandleFunc("POST " + updatePath, func(w http.ResponseWriter, r *http.Request) {
			updatePostHandler(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := indexTmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
