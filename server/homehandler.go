package server

import (
	"html/template"
	"net/http"
)



func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // only get method allowed
		ErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// If the list of artists is empty, it fetches the data.
	if len(Artist) == 0 {
		FetchData(w)
	}
	// Depending on the URL path, it either serves the home page template or returns a 404 error.
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
