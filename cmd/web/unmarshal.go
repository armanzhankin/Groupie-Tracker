package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func AllBands() error {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	jsonErr := json.Unmarshal(body, &PeopleNumArtists)
	if jsonErr != nil {
		log.Println("Can not unmarshall JSON")
		return jsonErr
	}

	return nil
}

func AllRelation() error {
	url := "https://groupietrackers.herokuapp.com/api/relation"

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	jsonErr := json.Unmarshal(body, &Relations)
	if jsonErr != nil {
		log.Println("Can not unmarshall JSON")
		return jsonErr
	}
	for i := range PeopleNumArtists {
		PeopleNumArtists[i].DatesLocations = Relations.Index[i].DatesLocations
		FormatMarks(i)
	}
	return nil
}
