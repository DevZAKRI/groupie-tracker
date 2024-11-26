package server

type Artists []struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

var Artist Artists

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
	Loc            []string
	ConDate        []string
}

type Error struct {
	Code    int
	Message string
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

var ArtLocations Location
var DatesData Dates
