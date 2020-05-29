package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strconv"
)

// handler for home
func (app *application) home(w http.ResponseWriter, r *http.Request) {

    // path check
    if r.URL.Path != "/" {
        app.notFound(w)
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
        app.serverError(w, err)
        return
    }

    // use execute to write template to response
    err = ts.Execute(w, nil)
    if err != nil {
        app.serverError(w, err)
    }

}

// handler for show snippet
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }

    fmt.Fprintf(w, "Display snippet with ID: %d...", id)
}

// handler for create snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

    w.Write([]byte("TODO: Create snippet"))
}
