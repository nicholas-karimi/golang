package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// _ "github.com/jinzhu/gorm/dialect/mysql"

	"github.com/nicholas-karimi/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
