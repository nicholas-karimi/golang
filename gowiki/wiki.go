package main

import (
	"fmt"
    "html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

/*func main(){
    p1 := &Page{Title: "TestPage", Body:[]byte("This is a sample page")}
    p1.save()

    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}
*/

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err //nil is the zero-value for poiners, interfaces etc
	}
	return &Page{Title: title, Body: body}, nil
}

// template handler
func renderTemplate(w http.ResponseWriter, templ string, p *Page){
    t,err := template.ParseFiles(templ + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
} 

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
	//	fmt.Println("Encounterd an error: ", err)
	    http.Redirect(w, r, "/edit/"+title, http.StatusFound)
    }
    //t, _ := template.ParseFiles("views.html")
    //t.Execute(w, p)
    renderTemplate(w, "views", p)
}

func editHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    //t, _ := template.ParseFiles("edit.html")
    //t.Execute(w, p)
    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/save"):]
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err = p.save()
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return

    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}


func main() {
	http.HandleFunc("/view/", viewHandler)		
   	http.HandleFunc("/edit/", editHandler)
    //http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
