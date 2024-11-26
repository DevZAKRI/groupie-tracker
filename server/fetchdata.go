package server

import (
	"encoding/json"
	"net/http"
	"sync"
)

// FetchData concurrently fetches data from multiple API endpoints and waits for all fetch operations to complete.
// The fetched data is stored in the respective global variables: Artist, Locations, DatesData, and Relation.
func FetchData(w http.ResponseWriter) {
	urlArtist := "https://groupietrackers.herokuapp.com/api/artists"
	//urlLoc := "https://groupietrackers.herokuapp.com/api/locations"
	//urlDate := "https://groupietrackers.herokuapp.com/api/dates"
	urlRelation := "https://groupietrackers.herokuapp.com/api/relation"

	var wg sync.WaitGroup
	wg.Add(2)
	go Fetch(w, urlArtist, &Artist, &wg)
	//go Fetch(w, urlLoc, &ArtLocations, &wg)
	//go Fetch(w, urlDate, &DatesData, &wg)
	go Fetch(w, urlRelation, &Relation, &wg)
	wg.Wait()
}

func Fetch(w http.ResponseWriter, url string, data interface{}, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		ErrorPage(w, "Internal Server Error", 500)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		ErrorPage(w, "Internal Server Error", 500)
		return
	}
	wg.Done()

}
