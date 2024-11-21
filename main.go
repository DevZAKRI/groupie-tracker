package main

import (
	"fmt"
	"groupie/server"
	"net/http"
)

func main() {
	//http.HandleFunc("/assets", server.AssetsHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", server.HomeHandler)
	fmt.Print("http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Server Failed!!")
	}

}

// func init() {
// 	var err error
// 	server.Tp, err = template.ParseGlob("./templates/*.html")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }
