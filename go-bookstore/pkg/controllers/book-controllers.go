package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/nicholas-karimi/go-bookstore/pkg/utils"
	"github.com/nicholas-karimi/go-bookstore/pkg/models"

)

var NewBook models.Book

func GetBook(w http.ResponseWriter, *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshall(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBookById(w http.ResponseWriter, *http.Request){
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(id, 0 ,0)
	if err != nil {
		fmt.Println("error while parsing:", err)
	}
	bookDetails, _ := models.GetBookByID(ID)
        res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}


func CreateBook(w http.ResponseWriter, *http.Request){
	CreateBook := &modle.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, *http.Request){
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)

}


func DeleteBook(w http.ResponseWriter, *http.Request){
	var updateBook = &model.Book{}
	utils.ParseBody(r, updatebook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {

		fmt.Println("error while parsing")
	
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	}

}



