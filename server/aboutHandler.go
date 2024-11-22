package server

import "net/http"

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/about.html")
}
