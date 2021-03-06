package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"path/filepath"
)

const storagePath = "./storage"
const staticPath = "./static"

// fun main()
func main() {

	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		os.MkdirAll(storagePath, os.ModePerm)
	}

	if _, err := os.Stat(filepath.Join(storagePath, staticPath)); os.IsNotExist(err) {
		os.MkdirAll(filepath.Join(storagePath, staticPath), os.ModePerm)
	}

	router := mux.NewRouter()

	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(filepath.Join(storagePath, staticPath)))))

	router.HandleFunc("{storage}/{id}", GetDocument).Methods("GET")

	router.HandleFunc("/api/v1/{storage}", CreateDocument).Methods("POST")
	router.HandleFunc("/api/v1/static/{client}/{folder}", CreateStaticFile).Methods("POST")

	log.Fatal(http.ListenAndServe(":8150", router))
}
