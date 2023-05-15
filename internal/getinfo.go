package groupie

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var artists []Artist

func GetArtists() ([]Artist, error) {
	AllArtists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer AllArtists.Body.Close()
	body, err := ioutil.ReadAll(AllArtists.Body)
	json.Unmarshal(body, &artists)
	return artists, nil
}

func GetOneArtist(i int) Artist {
	AllArtists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	defer AllArtists.Body.Close()
	body, err := ioutil.ReadAll(AllArtists.Body)
	json.Unmarshal(body, &artists)
	return artists[i-1]
}

func GetRelations(i int) map[string][]string {
	rel, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(i))
	if err != nil {
		log.Fatal(err)
	}
	defer rel.Body.Close()
	body, err := ioutil.ReadAll(rel.Body)
	if err != nil {
		log.Fatal(err)
	}
	var relat Relations
	json.Unmarshal(body, &relat)
	res := relat.LD
	return res
}
