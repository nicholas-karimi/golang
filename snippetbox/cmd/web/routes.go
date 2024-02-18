package routes

import (
	"net/http"

	"github.com/nicholas-karimi/snippetbox/cmd/web/handlers"
	"github.com/nicholas-karimi/snippetbox/config"
)

// The routes() method returns a servemux containing our application routes.

func routes(app *Application) *http.ServeMux {
	logger := config.NewLogger()

	mux := http.NewServeMux()

	// file server to serve files from the static directory
	fileServer := http.FileServer(http.Dir("ui/static"))

	// register the file server as the handler for the static directory
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// other app routes
	mux.HandleFunc("/", handlers.Home(logger))
	mux.HandleFunc("/snippet/view", handlers.SnippetView)
	mux.HandleFunc("/snippet/create", handlers.SnippetCreate)

	return mux
}
