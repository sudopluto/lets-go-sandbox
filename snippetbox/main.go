package main

import (
	"log"
	"net/http"
)

// handler for home
// write "hello" to response, ignore incoming request
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// create new router, and map home() to /
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	//start logging and start up webserver
	// listen and serve should never return unless hit error
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
