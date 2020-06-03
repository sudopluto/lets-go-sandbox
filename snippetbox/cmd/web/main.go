package main

import (
    "database/sql"
    "flag"
	"log"
    "net/http"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}

func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}


func main() {

    addr := flag.String("addr", ":4000", "HTTP network address")
    dsn := flag.String("dsn", "web:snippetbox@/snippetbox?parseTime=true",
        "MySQL data source name")

    flag.Parse()

    infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)

    errorLog := log.New(os.Stderr, "[ERR]\t",
        log.Ldate|log.Ltime|log.Lshortfile)

    db, err := openDB(*dsn)
    if err != nil {
        errorLog.Fatal(err)
    }

    defer db.Close()

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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
