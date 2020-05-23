package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
)

// handler for home
func home(w http.ResponseWriter, r *http.Request) {

    // path check
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    files := []string{
        "./ui/html/home.page.tmpl",
        "./ui/html/base.layout.tmpl",
        "./ui/html/footer.partial.tmpl",
    }

    // try to parse template, throw error if can't
    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    // use execute to write template to response
    err = ts.Execute(w, nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }

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
