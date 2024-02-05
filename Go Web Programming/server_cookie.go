package main

import(
    "net/http"
    "encoding/base64"
    "time"
    "fmt"
)

func setCookie(w http.ResponseWriter, r *http.Request){
    c1 := http.Cookie{
        Name: "first_cookie",
        Value: "Go Web Programming",
        HttpOnly: true,
    }

    c2 := http.Cookie{
         Name: "second_cookie",
        Value: "Manning Publication Co",
        HttpOnly: true,

    }
   // w.Header().Set("Set-Cookie", c1.String())
    // w.Header().Add("Set-Cookie", c2.String())
    
    //Go shortcut
    http.SetCookie(w, &c1)
    http.SetCookie(w, &c2)

}

// cookie for flash messafes
func setMessage(w http.ResponseWriter, r *http.Request){
    msg := []byte ("Hello World!")
    c := http.Cookie{
        Name: "flash",
        Value: base64.URLEncoding.EncodeToString(msg),
    }
    http.SetCookie(w, &c)
}
func showMessage(w http.ResponseWriter, r *http.Request){
    c, err := r.Cookie("flash")
    if err != nil {
        if err == http.ErrNoCookie {
            fmt.Fprintln(w, "No message found")
        }
    } else{
        rc := http.Cookie{
            Name: "flash",
            MaxAge: -1,
            Expires: time.Unix(1, 0),

        }
        http.SetCookie(w, &rc)
        val, _ := base64.URLEncoding.DecodeString(c.Value)
        fmt.Fprintln(w, string(val))
    }
}


func getCookie(w http.ResponseWriter, r *http.Request){
    h := r.Header["Cookie"]
    fmt.Fprintln(w, h)
}
func main(){

    server := http.Server{
        Addr: "127.0.0.1:8080",
    }

    http.HandleFunc("/set_cookie", setCookie)
    http.HandleFunc("/get_cookie", getCookie)
    
    http.HandleFunc("/set_message", setMessage)
    http.HandleFunc("/show_message", showMessage)
    
    server.ListenAndServe()

}
