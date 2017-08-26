package routes

import (
	"os"
	"fmt"
	"strings"
	"net/http"
	"path/filepath"
	"github.com/go-chi/chi"
	"github.com/robert-hansen/goapp/controller"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "public")
	FileServer(router, "/static", http.Dir(filesDir))

	router.Get("/", Index)
	router.Get("/hello/:name", Hello)

	return router
}

func Index(res http.ResponseWriter, req *http.Request) {
	controller.Home(res, req)
}

func Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello ", req.Context().Value("name"))
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}