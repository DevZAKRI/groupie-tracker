package server

import (
	"net/http"
	"strings"
)

func AssetsHandler(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/assets/") {
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
