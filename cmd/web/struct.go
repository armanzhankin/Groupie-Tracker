package cmd

import "net/http"

const (
	HomePage     string = "ui/html/index.html"
	ArtistPage   string = "ui/html/artist.html"
	ErrorPage    string = "ui/html/error.html"
	TemplatePage string = "ui/html/templates.html"
	NeuteredPage string = "ui/neutered.html"
)

type NeuteredFileSystem struct {
	fs http.FileSystem
}

type Err struct {
	Status int
	Text   string
}

type Artists struct {
	Id             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	Locations      string              `json:"locations"`
	ConcertDates   string              `json:"concertDates"`
	FirstAlbum     string              `json:"firstalbum"`
	Relations      string              `json:"relations"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relation struct {
	Index []struct {
		Id             uint64              `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}
}

var (
	PeopleNumArtists []Artists
	Relations        Relation
)
