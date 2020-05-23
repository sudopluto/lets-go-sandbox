package main

import (
	"log"
    "net/http"
)


func main() {

	// create new router
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

    fileServer := http.FileServer(http.Dir("./ui/static/"))

    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// start logging and start up webserver
	// listen and serve should never return unless hit error
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
