package main

import (
	"github.com/gofrs/uuid"
	"os"
	"path/filepath"
	"regexp"
)

func GetAndCreateStoragePath(storage string) (string, uuid.UUID, error) {

	u := uuid.Must(uuid.NewV4())

	imagePath := GetPath(storage, u.String())

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		os.MkdirAll(imagePath, os.ModePerm)
	}

	return imagePath, u, nil
}

func GetPath(storage string, uid string) (string) {
	return filepath.Join(storagePath, storage, uid[0:2], uid[2:4], uid)
}

func CreateFileName(fileName string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(fileName, ""), nil
}
