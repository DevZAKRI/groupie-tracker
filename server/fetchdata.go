package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func FetchData() {
	urlArtist := "https://groupietrackers.herokuapp.com/api/artists"
	urlLoc := "https://groupietrackers.herokuapp.com/api/locations"
	urlDate := "https://groupietrackers.herokuapp.com/api/dates"
	urlRelation := "https://groupietrackers.herokuapp.com/api/relation"

	var wg sync.WaitGroup
	wg.Add(4)
	go Fetch(urlArtist, &Artist, &wg)
	go Fetch(urlLoc, &Locations, &wg)
	go Fetch(urlDate, &DatesData, &wg)
	go Fetch(urlRelation, &Relation, &wg)
	wg.Wait()
}

func Fetch(url string, data interface{}, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data")
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		fmt.Println("Error decoding data")
		return
	}
	wg.Done()

}
