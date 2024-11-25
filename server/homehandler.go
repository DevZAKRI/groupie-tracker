package server

import (
	"html/template"
	"net/http"
)

var Artist Artists

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	FetchData()
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
