package main

import (
    "flag"
	"log"
    "net/http"
    "os"
)

type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}


func main() {

    addr := flag.String("addr", ":4000", "HTTP network address")

    flag.Parse()

    infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)

    errorLog := log.New(os.Stderr, "[ERR]\t",
        log.Ldate|log.Ltime|log.Lshortfile)

    app := &application{
        errorLog: errorLog,
        infoLog:  infoLog,
    }

	// create new router
    mux := app.routes()

    srv := &http.Server{
        Addr: *addr,
        ErrorLog: errorLog,
        Handler: mux,
    }

	// start logging and start up webserver
	// listen and serve should never return unless hit error
	infoLog.Printf("Starting server on %s", srv.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
