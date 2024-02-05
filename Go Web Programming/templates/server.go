package main

import (
    "net/http"
    "html/template"
    "math/rand"
    "time"
    "fmt"
    "log"
)

func process (w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("tmpl.html")
    
    if err != nil {
        fmt.Println("An error occured: ", err)
    }
    rand.Seed(time.Now().Unix())

    t.Execute(w, rand.Intn(10) > 5)
    daysOfWeek := [] string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
    t.Execute(w, daysOfWeek)
}

func tempInclude(w http.ResponseWriter, r *http.Request){
    t, err := template.ParseFiles("t1.html", "t2.html")
    if err != nil {
        fmt.Println("An error occured here: ", err)

    }
    t.Execute(w, "Hello Karimi!")
}

// custom template functions
func formatDate(t time.Time) string {
    layout := "2006-01-02" //ISO8601
    return t.Format(layout)

}

func templateFunctions(w http.ResponseWriter, r *http.Request){
    funcMap := template.FuncMap {"fdate": formatDate }
    t := template.New("test.html").Funcs(funcMap)
    t, _ = t.ParseFiles("test.html")

    t.Execute(w, time.Now())

}
func main(){
   
    http.HandleFunc("/process", process)
    http.HandleFunc("/temps", tempInclude)
    http.HandleFunc("/time", templateFunctions)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

