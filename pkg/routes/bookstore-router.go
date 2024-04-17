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

	router.HandleFunc("POST /author/", controllers.CreateAuthor)
	router.HandleFunc("GET /authors/", controllers.GetAuthors)
	router.HandleFunc("DELETE /author/{id}", controllers.DeleteAuthor)
	router.HandleFunc("GET /author/{id}/", controllers.GetAuthorById)
	router.HandleFunc("PUT /author/{id}/", controllers.UpdateAuthor)

	router.HandleFunc("GET /login/google/", controllers.GoogleLogin)
	router.HandleFunc("GET /login/google/callback/", controllers.GoogleCallback)
}
