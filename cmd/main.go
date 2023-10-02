package main

import (
	"log"
	"net/http"

	routes "github.com/RijinSwaminathan/book-management-system/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.HandleFunc("/", r)
	log.Fatal(http.ListenAndServe(":8010", nil))
}
