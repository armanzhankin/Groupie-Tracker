package cmd

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(url[2])
	if r.URL.Path != "/artist/"+url[2] {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if err != nil || len(url) > 3 || id > len(PeopleNumArtists) || id < 1 {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles(ArtistPage, TemplatePage)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}

	err = AllRelation()
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	// if len(PeopleNumArtists) == 0 {
	// 	err = AllBands()
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		ErrorHandler(w, http.StatusInternalServerError)
	// 		return
	// 	}
	// }

	err = tmpl.Execute(w, PeopleNumArtists[id-1])
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
