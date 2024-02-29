package routes

import (
	"net/http"

	"github.com/maratb3k/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *http.ServeMux) {
	router.HandleFunc("POST /book/", controllers.CreateBook)
	router.HandleFunc("GET /books/", controllers.GetBooks)
	router.HandleFunc("DELETE /book/{id}", controllers.DeleteBook)
	router.HandleFunc("GET /book/{id}/", controllers.GetBookById)
	router.HandleFunc("PUT /book/{id}/", controllers.UpdateBook)
	/*router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBookById).Methods("GET")
	router.HandleFunc("/books", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")*/
}
