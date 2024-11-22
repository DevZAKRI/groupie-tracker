package server

import (
	"fmt"
	"html/template"
	"net/http"
)

var Artist Artists

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	tmp.Execute(w, Artist)

}
