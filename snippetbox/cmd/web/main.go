package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/nicholas-karimi/snippetbox/cmd/web/handlers"
)

func main() {
	// command line flag - default value :4000 & help  text to expalin what the flag controls
	addr := flag.String("addr", ":4000", "HTTP network address")

	// parse the commabndline flag
	flag.Parse()

	// use log.New() to create a logger for writing information messages.
	// takes 3 params: destination to write logs(os.Stdout), string prefix for message(INFO)
	//and tags for additional ifno to include (local date and time - flags joined using bitwise OR operator |)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// error msg logger - use stderr as destination and log.Lshortfile flag to include relevant filebame and line number
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	// file server to serve files from the static directory
	fileServer := http.FileServer(http.Dir("ui/static"))

	// register the file server as the handler for the static directory
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// other app routes
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/snippet/view", handlers.SnippetView)
	mux.HandleFunc("/snippet/create", handlers.SnippetCreate)

	// log.Print("Starting web server on :4000")
	// log.Printf("Starting server on %s", *addr)
	infoLog.Printf("Starting server on %s", *addr)

	// err := http.ListenAndServe(":4000", mux)
	err := http.ListenAndServe(*addr, mux)

	// log.Fatal(err)
	errLog.Fatal(err)
}
