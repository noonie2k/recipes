package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var data map[string]*Item

func main() {
	data = ReadData()

	fmt.Println(GetEnvPort())

	router := NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./app/")))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", GetEnvPort()), router))
}

func GetEnvPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return port
}
