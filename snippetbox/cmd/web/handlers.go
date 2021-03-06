package main

import (
    "fmt"
//    "html/template"
    "net/http"
    "strconv"

    "github.com/sudopluto/lets-go-sandbox/pkg/models"
)

// handler for home
func (app *application) home(w http.ResponseWriter, r *http.Request) {

    // path check
    if r.URL.Path != "/" {
        app.notFound(w)
        return
    }

    s, err := app.snippets.Latest()
    if err != nil {
        app.serverError(w, err)
        return
    }

    for _, snippet := range s {
        fmt.Fprintf(w, "%v\n", snippet)
    }

    /*
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
    */

}

// handler for show snippet
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        app.notFound(w)
        return
    }

    s, err := app.snippets.Get(id)
    if err == models.ErrNoRecord {
        app.notFound(w)
        return
    } else if err != nil {
        app.serverError(w, err)
        return
    }

    fmt.Fprintf(w, "%v", s)
}

// handler for create snippet
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        app.clientError(w, http.StatusMethodNotAllowed)
        return
    }

    title := "O snail"
    content := "O snail\nClimb Mount Fuji\nBut slowly, slowly!\n\n - Kobayashi"
    expires := "7"

    id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, err)
        return
    }

    http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)

    w.Write([]byte("TODO: Create snippet\n"))
}
