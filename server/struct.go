package server

// var Tp *template.Template
var URL = "https://groupietrackers.herokuapp.com/api/artists"

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	//Locations    string   `json:"locations"`
	//ConcertDates string   `json:"concertDates"`
	Relations string `json:"relations"`
}

type Relations struct {
	Id             int
	LocationsDates map[string][]string
}
