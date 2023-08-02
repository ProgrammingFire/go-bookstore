package routes

import (
	"github.com/gorilla/mux"
	"github.com/programmingfire/bookstore/pkg/controllers"
)

func HandleRoutes(r *mux.Router) {
	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	r.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}
