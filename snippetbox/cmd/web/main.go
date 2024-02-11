package main

import (
	"log"
	"net/http"
	"flag"

	"github.com/nicholas-karimi/snippetbox/cmd/web/handlers"
)

func main() {
	// command line flag - default value :4000 & help  text to expalin what the flag controls
	addr := flag.String("addr", ":4000", "HTTP network address")

	// parse the commabndline flag
	flag.Parse()

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
	log.Printf("Starting server on %s", *addr)

	// err := http.ListenAndServe(":4000", mux)
	err := http.ListenAndServe(*addr, mux)
	
	log.Fatal(err)
}
