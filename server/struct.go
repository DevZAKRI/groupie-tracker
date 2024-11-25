package server

// var Tp *template.Template
var URL = "https://groupietrackers.herokuapp.com/api/artists"
var URL1 = "https://groupietrackers.herokuapp.com/api/relation"

type Artists []struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	//Locations    string   `json:"locations"`
	//ConcertDates string   `json:"concertDates"`
	//Relations    string   `json:"relations"`
}

type RelationD struct {
	Id             int                 `json:"id"`
	LocationsDates map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Index []RelationD `json:"index"`
}

var Relation Relations

var ArtistData struct {
	Id             int
	Image          string
	Name           string
	Members        []string
	CreationDate   int
	FirstAlbum     string
	LocationsDates map[string][]string
}

type Error struct {
	Code    int
	Message string
}

type Location struct {
	Locations []string `json:"locations"`
}

type Dates struct {
	Dates []string `json:"dates"`
}

var Locations Location
var DatesData Dates
