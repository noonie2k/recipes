package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var data map[string]*Item
var dataList []*Item

func main() {
	data = ReadData()

	for itemId, item := range data {
		item.Identifier = itemId
		dataList = append(dataList, item)
	}

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
