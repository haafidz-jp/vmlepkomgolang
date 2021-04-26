package main

import (
	"log"
	"net/http"
	"repo/4/handler" //sesuaikan dengan nama folder (case sensitive)
)

func main() {
	http.HandleFunc("/api/", handler.API)
	//Ganti 2 digit akhir port dengan 2 digit akhir NPM anda
	log.Println("localhost : 8087")
	http.ListenAndServe(":8087", nil)
}
