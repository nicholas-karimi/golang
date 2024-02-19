package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/nicholas-karimi/snippetbox/config"
)

// Handlers struct holds application-wide dependencies
type Handlers struct {
	AppConfig *config.AppConfig
}

// func Home(app *config.Application) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
	func(h *Handlers) Home(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			// app.notFound(w)
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
			h.AppConfig.InfoLog.Fatal(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		// err = ts.Execute(w, nil)
		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			// log.Print(err.Error())
			// app.LogError(err.Error())
			h.AppConfig.ErrorLog.Fatal(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}

func (h *Handlers) SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func (h *Handlers) SnippetCreate(w http.ResponseWriter, r *http.Request) {
	// snippet model
	snippetModel := h.AppConfig.SnippetModel

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// w.Write([]byte("Create a new snippet"))
	// dummy data
	title := "I am who I think I am."
	content := "Listen up, I am Karimi, an exception Python and Golang programmer. Do you know why I think you thing I am lonely. Its because you can't apprecaite the fact I have grown. To be exception, who have to do what the average person is too afraid to do."

	expires := 7

	id, err := snippetModel.Insert(title, content, time.Now(), expires)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError) //500)
		h.AppConfig.ErrorLog.Fatal(err.Error())
		return
	}
	fmt.Fprintf(w, "Snippet successfully created with ID %d", id)

	// redirect user to relevant pagefor the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
