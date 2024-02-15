package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/nicholas-karimi/snippetbox/config"
)

func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/* dependencyn injection
		app := r.Context() .Value("app").(*application)*/
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		// w.Write([]byte("Hello from Snippetbox!"))
		files := []string{
			"ui/html/base.tmpl",
			"ui/html/partials/nav.tmpl",
			"ui/html/pages/home.tmpl",
		}
		// ts, err := template.ParseFiles("ui/html/pages/home.tmpl")
		ts, err := template.ParseFiles(files...)
		if err != nil {
			// log.Print(err.Error())
			// http.Error(w, "Internal Server Error", 500)
			app.LogError(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		// err = ts.Execute(w, nil)
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			// log.Print(err.Error())
			app.LogError(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
}
func SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))

}
