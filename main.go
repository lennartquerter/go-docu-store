package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"
)

const storagePath = "./storage"
const staticPath = "./static"

// fun main()
func main() {

	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		os.MkdirAll(storagePath, os.ModePerm)
	}

	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		os.MkdirAll(staticPath, os.ModePerm)
	}

	router := mux.NewRouter()

	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(staticPath))))

	router.HandleFunc("/{storage}", CreateDocument).Methods("POST")
	router.HandleFunc("/{storage}/{id}", GetDocument).Methods("GET")
	router.HandleFunc("/{client}/{folder}", CreateStaticFile).Methods("POST")

	log.Fatal(http.ListenAndServe(":8150", router))
}
