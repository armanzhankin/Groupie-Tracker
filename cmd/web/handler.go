package cmd

import (
	"fmt"
	"log"
	"net/http"
)

func Handler() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(NeuteredFileSystem{http.Dir("ui")})
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")

	mux.Handle("/ui", http.NotFoundHandler())
	mux.Handle("/ui/", http.StripPrefix("/ui/", fileServer))

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/artist/", Artist)

	// mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir("ui"))))
	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal("ERROR:Server not listening")
	}
}
