package main

import (
	"fmt"
	"net/http"
)

func getHeaders(w http.ResponseWriter, r *http.Request) {
	//h := r.Header

	// get specif /header
	// h2 := r.Header.Get("Accept-Encoding") // gzip, deflate, br
	h2 := r.Header["Accept-Encoding"] // [gzip, deflate, br]
	fmt.Fprintln(w, h2)
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", getHeaders)

	//log.Fatal(http.ListenAndServe(":8080", nil))
	server.ListenAndServe()
}
