package server

import (
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	idStr := r.URL.Path[8:] // "/artist/12"
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorPage(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if len(Artist) == 0 {
		FetchData()
	}
	for _, oneGroup := range Artist {
		if oneGroup.Id == id {
			ArtistData.Id = oneGroup.Id
			ArtistData.Image = oneGroup.Image
			ArtistData.Name = oneGroup.Name
			ArtistData.Members = oneGroup.Members
			ArtistData.CreationDate = oneGroup.CreationDate
			ArtistData.FirstAlbum = oneGroup.FirstAlbum
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
		ErrorPage(w, "Page Not Found", http.StatusNotFound)
		return
	}

}
