package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 9000

	http.Handle("/", http.FileServer(http.Dir("polymer")))
	fmt.Printf("Server Berjalan pada port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
