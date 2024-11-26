package server

import (
	"html/template"
	"net/http"
)

var Artist Artists

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	if len(Artist) == 0 {
		FetchData()
	}
	switch r.URL.Path {
	case "/":
		tmp, err := template.ParseFiles("templates/home.html")
		if err != nil {
			ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmp.Execute(w, Artist)
	default:
		ErrorPage(w, "Page Not Found", http.StatusNotFound)
	}
}
