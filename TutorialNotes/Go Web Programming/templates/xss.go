package main

import (
	"html/template"
	"net/http"
)

func processXss(w http.ResponseWriter, r *http.Request) {
    // stop browser from xss protection
    w.Header().Set("X-XSS-Protection", "0")
	t, _ := template.ParseFiles("t.html")
	// t.Execute(w, r.FormValue("comment"))
	t.Execute(w, template.HTML(r.FormValue("comment"))) // unescape special chars to allow xss

}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)

}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process_xss", processXss)
	http.HandleFunc("/form", form)

	server.ListenAndServe()
}
