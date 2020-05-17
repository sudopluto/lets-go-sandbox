package main

import (
    "fmt"
	"log"
    "net/http"
    "strconv"
)

// handler for home
func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

	w.Write([]byte("Hello from Snippetbox"))
}

// handler for show snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Display snippet with ID: %d...", id)
}

// handler for create snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        w.WriteHeader(405)
        w.Write([]byte("Method not allowed"))
        return
    }

    w.Write([]byte("TODO: Create snippet"))
}

func main() {

	// create new router
    // - home() to /
    // - showSnippet() to /snippet
    // - createSnippet() to /snippet/create
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//start logging and start up webserver
	// listen and serve should never return unless hit error
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
