package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	Data, err := http.Get(URL)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer Data.Body.Close()

	var artists Artist
	err = json.NewDecoder(Data.Body).Decode(&artists)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "Error GO Away")
		return
	}

	tmp.Execute(w, artists)
}
