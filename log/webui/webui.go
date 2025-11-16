package webui

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"net/url"
	"syscall"
)

const (
	basePath = "/log/level"
)

type NoDirFileSystem struct {
	fs http.FileSystem
}

func FsNoDirFileServer(fsys fs.FS) http.Handler {
	return NoDirFileServer(http.FS(fsys))
}

func NoDirFileServer(fs http.FileSystem) http.Handler {
	return http.FileServer(NewNoDirFileSystem(fs))
}

func (fs NoDirFileSystem) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
		return nil, syscall.ENOENT
	}

	return f, nil
}

func NewNoDirFileSystem(fs http.FileSystem) NoDirFileSystem {
	return NoDirFileSystem{fs: fs}
}

var (
	//go:embed templates/*.tmpl
	TemplatesFS embed.FS

	//go:embed all:static
	StaticFS       embed.FS
	staticSubFS, _ = fs.Sub(StaticFS, "static") // strip "static" prefix from FS
	staticHandler  = FsNoDirFileServer(staticSubFS)

	indexTmpl = template.Must(template.ParseFS(TemplatesFS, "templates/index.tmpl"))
)

// Mount attaches all handlers to a mux.
func Mount(mux *http.ServeMux) {
	// Statics
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	// Templates
	mux.HandleFunc("GET "+basePath, indexHandler)

	// RPC
	levelPath, _ := url.JoinPath(basePath, "all")
	mux.HandleFunc("GET "+levelPath, levelHandler)

	levelPackagePath, _ := url.JoinPath(basePath, "package", "all")
	mux.HandleFunc("GET "+levelPackagePath, packageHandler)

	updatePath, _ := url.JoinPath(basePath, "package", "update")
	mux.HandleFunc("POST "+updatePath, func(w http.ResponseWriter, r *http.Request) {
		updatePostHandler(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := indexTmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
