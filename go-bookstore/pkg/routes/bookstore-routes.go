package routes

import (
	"github.com/gorilla/mux"
	"github.com/nicholas-karimi/go-bookstore/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router){ 
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")	
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
