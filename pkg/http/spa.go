package http

import (
	"errors"
	"io/fs"
	"net/http"
	"path/filepath"
	"time"
)

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	static    http.FileSystem
	indexPath string
}

func NewSPAHandler(static http.FileSystem, indexPath string) http.Handler {
	return &spaHandler{
		static:    static,
		indexPath: indexPath,
	}
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check whether a file exists at the given path
	_, err = h.static.Open(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			// file does not exist, serve index.html
			indexFile, err := h.static.Open(h.indexPath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.ServeContent(w, r, h.indexPath, time.Time{}, indexFile)
			return
		}

		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(h.static).ServeHTTP(w, r)
}
