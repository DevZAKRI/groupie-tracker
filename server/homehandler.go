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
	artistsData, err := http.Get(URL)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer artistsData.Body.Close()
	var artist []Artist
	err1 := json.NewDecoder(artistsData.Body).Decode(&artist)
	if err1 != nil {
		fmt.Println(err1)
		fmt.Fprint(w, "Error GO Away")
	}

	fmt.Println(artistsData.Body)
	tmp.Execute(w, artist)
}
