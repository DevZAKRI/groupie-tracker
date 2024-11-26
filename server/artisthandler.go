package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { // only get method allowed
		ErrorPage(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorPage(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// find the artist ID if it is invalid, it returns a "Page Not Found" error.
	idStr := r.URL.Path[8:] // "/artist/12"
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorPage(w, "Page Not Found", http.StatusNotFound)
		return
	}
	// if the artist data is not already fetched, it calls FetchData to retrieve it.
	if len(Artist) == 0 {
		FetchData(w)
	}
	// It then searches for the artist by ID in the Artist slice and populates the ArtistData struct.
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
					break
				}
			}
			err := fetch("https://groupietrackers.herokuapp.com/api/locations/"+idStr, &ArtLocations)
			if err != nil {
				fmt.Println(err)
			}
			ArtistData.Loc = ArtLocations.Locations
			err = fetch("https://groupietrackers.herokuapp.com/api/dates/"+idStr, &DatesData)
			if err != nil {
				fmt.Println(err)
			}
			ArtistData.ConDate = DatesData.Dates

		}
		//fmt.Println(ArtistData)
	}
	// If the artist is found, it executes the template with the artist data.
	// If the artist is not found, it returns a "Page Not Found" error.
	if len(Artist) >= id {
		tmp.Execute(w, ArtistData)
	} else {
		ErrorPage(w, "Page Not Found", http.StatusNotFound)
		return
	}
}

func fetch(erl string, data interface{}) error {
	resp, err := http.Get(erl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}
