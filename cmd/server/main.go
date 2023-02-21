package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":4000", http.FileServer(http.Dir("../../assets"))))
}
