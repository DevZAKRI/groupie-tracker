package server

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	idStr := r.URL.Path[8:]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	for _, h := range Artist {
		if h.Id == id {
			ArtistData.Id = h.Id
			ArtistData.Image = h.Image
			ArtistData.Name = h.Name
			ArtistData.Members = h.Members
			ArtistData.CreationDate = h.CreationDate
			ArtistData.FirstAlbum = h.FirstAlbum
			for _, h := range Relation.Index {

				if h.Id == id {
					ArtistData.LocationsDates = h.LocationsDates
				}
			}
		}
	}

	if len(Artist) >= id {

		tmp.Execute(w, ArtistData)
	} else {
		http.NotFound(w, r)
		return
	}

}
