package main

import (
	"fmt"
	groupie "groupie-tracker/internal"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type DataErr struct {
	Header  int
	Message string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrExec(w, http.StatusNotFound, "Page is not found")
		return
	}

	if r.Method != http.MethodGet {
		ErrExec(w, http.StatusMethodNotAllowed, "Method is not allowed")
		return
	}

	ts, err := template.ParseFiles("../ui/html/home.html")
	if err != nil {
		log.Println(err.Error())
		ErrExec(w, http.StatusInternalServerError, "Internal server error")

		return
	}

	data, err := groupie.GetArtists()
	if err != nil {
		log.Println(err.Error())
		ErrExec(w, http.StatusInternalServerError, "Internal server error")

		return
	}
	err = ts.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		ErrExec(w, http.StatusInternalServerError, "Internal server error")

	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		fmt.Println("query")
		ErrExec(w, http.StatusNotFound, "Page is not found")
		return
	}
	if id > 52 {
		fmt.Println("id >")
		ErrExec(w, http.StatusNotFound, "Page is not found")
		return
	}

	if r.Method != http.MethodGet {
		ErrExec(w, http.StatusMethodNotAllowed, "Method is not allowed")
		return
	}

	ts, err := template.ParseFiles("../ui/html/artist.html")
	if err != nil {
		ErrExec(w, http.StatusInternalServerError, "Intenal Server Error")
		log.Println(err.Error())
		return
	}

	arts := groupie.GetOneArtist(id)
	arts.LocationDate = groupie.GetRelations(id)
	if err != nil {
		log.Println(err.Error())
	}

	err = ts.Execute(w, arts)
	if err != nil {
		log.Println(err.Error())
		ErrExec(w, 500, "Intenal Server Error")
	}
}

func ErrExec(r http.ResponseWriter, header int, message string) {
	r.WriteHeader(header)
	tempErr, err := template.ParseFiles("../ui/html/error.html")
	if err != nil {
		http.Error(r, "", header)
		log.Print(err)
		return
	}
	data := DataErr{
		Header:  header,
		Message: message,
	}
	tempErr.ExecuteTemplate(r, "error", data)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
