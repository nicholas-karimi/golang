package main

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "net/http"
)


type HelloHandler struct{}

type WorldHandler struct{}

func (h *HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    fmt.Println(w, "Hello")
}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
    fmt.Println(w, "World!")
}


// read data from a request body
func body(w http.ResponseWriter, r *http.Request){
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    fmt.Fprintln(w, string(body))
}

func headers(w http.ResponseWriter, r *http.Request){
    h := r.Header
    get_h := r.Header.Get("Accept-Encoding")
    fmt.Fprintln(w, h, get_h)
}

// forms
func process (w http.ResponseWriter, r *http.Request){
    //r.ParseForm()
    r.ParseMultipartForm(1024)
    fileHeader := r.MultipartForm.File["uploaded"][0]
    file, err := fileHeader.Open()

    if err == nil {
        data, err := ioutil.ReadAll(file)
        if err == nil {
            fmt.Fprintln(w, string(data))
            }  

    }

}

type  Post struct {
    User string
    Threads [] string
}

func jsonExample(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    post := &Post{
        User: "Nicholas Karimi",
        Threads: []string{"first", "second", "third", "morning"},
    }
    json, _ := json.Marshal(post)
    w.Write(json)
}
func main(){
    hello := HelloHandler{}
    world := WorldHandler{}
    
    server := http.Server{
        Addr: "localhost:8080",
    }

    http.Handle("/hello", &hello)
    http.Handle("world", &world)
    
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/body", body)
    http.HandleFunc("/process", process)
    http.HandleFunc("/json", jsonExample)

    server.ListenAndServe()
}

