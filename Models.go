package main

type Image struct {
	ID        string `json:"Id,omitempty"`
	ImageName string `json:"ImageName,omitempty"`
	File      []byte
}


type ImageResponse struct {
	DocumentId string `json:"DocumentId"`
}