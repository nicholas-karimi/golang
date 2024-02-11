package main

import (
	"log"
	"net/http"

	"github.com/nicholas-karimi/snippetbox/cmd/web/handlers"
)

func main() {
	mux := http.NewServeMux()

	// file server to serve files from the static directory
	fileServer := http.FileServer(http.Dir("ui/static"))

	// register the file server as the handler for the static directory
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// other app routes
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/snippet/view", handlers.SnippetView)
	mux.HandleFunc("/snippet/create", handlers.SnippetCreate)

	log.Print("Starting web server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
