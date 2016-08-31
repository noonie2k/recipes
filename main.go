package main

import (
	"log"
	"net/http"
)

var data map[string]*Item

func main() {
	data = ReadData()

	router := NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./app/")))

	log.Fatal(http.ListenAndServe(":8080", router))
}
