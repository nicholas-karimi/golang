package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	// Year     int
	Year     string
	Director string
}

func main() {
	log.Println("Starting app...")
	fmt.Println("Hello, World!")

	/* root url */
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		// io.WriteString(w, "Hello, World!\n")
		// io.WriteString(w, r.Method)
		// tmpl.Execute(w, nil)
		films := map[string][]Film{
			"Films": {
				{Title: "The Matrix", Year: "1999", Director: "Wachowski"},
				{Title: "The Matrix Reloaded", Year: "2003", Director: "Wachowski"},
				{Title: "The Matrix Revolutions", Year: "2003", Director: "Wachowski"},
			},
		}
		tmpl.Execute(w, films)
	}
	// add film url
	h2 := func(w http.ResponseWriter, r *http.Request) {
		// simulate latency to test hx-indicator
		time.Sleep(5 * time.Second)
		// tmpl := template.Must(template.ParseFiles("add-film.html"))
		// tmpl.Execute(w, nil)
		// log.Print("HTMX Request from: ", r.RemoteAddr)
		// log.Print(r.Header.Get("HX-Request"))
		title := r.PostFormValue("title")
		year := r.PostFormValue("year")
		director := r.PostFormValue("director")
		// fmt.Println(title, year, director)
		/* Manual way to render a fragment in go
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-secondary text-white'>%s- %s - %s</li>", title, year, director)
		// fmt.Fprintf(w, htmlStr)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil) 
		*/
		templ := template.Must(template.ParseFiles("index.html"))
		templ.ExecuteTemplate(w, "film-list-item", Film{Title: title, Year: year, Director: director})
		}


	// url handler
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
