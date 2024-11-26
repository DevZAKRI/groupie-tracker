package main

import (
	"fmt"
	"groupie/server"
	"net/http"
)

func main() {
	http.HandleFunc("/assets/", server.AssetsHandler)
	http.HandleFunc("/artist/", server.ArtistHandler)
	http.HandleFunc("/", server.HomeHandler)
	http.HandleFunc("/about/", server.AboutHandler)
	fmt.Print("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Server Failed!!")
	}

}
