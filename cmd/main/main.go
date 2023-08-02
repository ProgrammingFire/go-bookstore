package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/programmingfire/bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	routes.HandleRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
