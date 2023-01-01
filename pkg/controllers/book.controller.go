package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"bookstore-go/pkg/models"
	"bookstore-go/pkg/utils"

	"github.com/gorilla/mux"
)

var Books models.Book

// index page
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Print("Hello Page")
}

// get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Println("Get books")
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Write(res)
}

// get book by id
func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	fmt.Println("Get book by id")

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// create book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	BookModel := &models.Book{}
	fmt.Println("Create book")

	utils.ParseBody(r, BookModel)
	b := BookModel.CreateBook()
	book, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(book)
}

// update book by id
func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Update book by id")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	BookModel := &models.Book{}
	utils.ParseBody(r, BookModel)
	b := BookModel.UpdateBook(ID)
	book, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(book)
}

// delete book by id
func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Delete book by id")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBookById(ID)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
