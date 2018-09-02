package main

import (
	"net/http"
	"path/filepath"
	"os"
	"io"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)

func CreateStaticFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	in, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer in.Close()

	var extension = filepath.Ext(header.Filename)
	var name = header.Filename[0 : len(header.Filename)-len(extension)]

	fileName, err := CreateFileName(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path := filepath.Join(staticPath, params["client"], params["folder"])

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	imagePath := filepath.Join(path, fileName+extension)

	out, err := os.OpenFile(imagePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer out.Close()
	io.Copy(out, in)

	resp := ImageResponse{
		DocumentId: "/" + imagePath,
	}

	json.NewEncoder(w).Encode(resp)
}

func CreateDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	in, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer in.Close()

	imagePath, uid, err := GetAndCreateStoragePath(params["storage"])

	fileName, err := CreateFileName(header.Filename)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path := filepath.Join(imagePath, fileName)

	out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer out.Close()
	io.Copy(out, in)

	resp := ImageResponse{
		DocumentId: uid.String(),
	}

	json.NewEncoder(w).Encode(resp)
}

func GetDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	path := GetPath(params["storage"], params["id"])

	files, err := ioutil.ReadDir(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(files) != 1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.ServeFile(w, r, filepath.Join(path, files[0].Name()))
}
