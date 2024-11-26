package server

import (
	"net/http"
	"strings"
)

// AssetsHandler handles HTTP requests for serving static assets.
// It only allows GET requests and serves files from the "./assets/" directory.
// If the request method is POST, it responds with "Method Not Allowed".
// If the requested asset path is empty or "/", it responds with "Access Forbidden".
// If the requested asset does not exist, it responds with "Page Not Found".
// Otherwise, it serves the requested asset file.
func AssetsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if strings.Contains(r.URL.Path, "/assets/") {
		filePath := strings.TrimPrefix(r.URL.Path, "/assets/")
		if filePath == "" || filePath == "/" {
			ErrorPage(w, "Access Forbidden", http.StatusForbidden)
			return
		}
		assetPath := "./assets/" + filePath
		err := AssetInfo(filePath)
		if err != "" {
			ErrorPage(w, "Page Not Found", http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, assetPath)
	} else {
		ErrorPage(w, "Page Not Found", http.StatusNotFound)
	}
}

func AssetInfo(name string) string {
	_, err := http.Dir("assets").Open(name)
	if err != nil {
		return "error"
	}
	return ""
}
