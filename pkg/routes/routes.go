package routes

import (
	"bookstore-go/pkg/controllers"

	"github.com/gorilla/mux"
)

// routes
var RegisterBookstoreRoutes = func(router *mux.Router) {
	// home page
	router.HandleFunc("/", controllers.Index).Methods("GET")
	// create book
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	// get all books
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	// get a book by id
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
	// delete a book
	router.HandleFunc("/books/{id}", controllers.DeleteBookById).Methods("DELETE")
	// update a book
	router.HandleFunc("/books/{id}", controllers.UpdateBookById).Methods("PUT")
}
