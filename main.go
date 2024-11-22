package main

import (
	"encoding/json"
	"fmt"
	"groupie/server"
	"net/http"
)

func main() {
	http.HandleFunc("/assets/", server.AssetsHandler)
	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/artist/", server.ArtistHandler)
	http.HandleFunc("/", server.HomeHandler)
	http.HandleFunc("/about/", server.AboutHandler)
	fmt.Print("http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Server Failed!!")
	}

}

func init() {
	artistData, err := http.Get(server.URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer artistData.Body.Close()
	relationsData, err := http.Get(server.URL1)
	if err != nil {
		fmt.Println(err)
	}
	defer relationsData.Body.Close()

	err = json.NewDecoder(artistData.Body).Decode(&server.Artist)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.NewDecoder(relationsData.Body).Decode(&server.Relation)
	if err != nil {
		fmt.Println(err)
		return
	}

}
