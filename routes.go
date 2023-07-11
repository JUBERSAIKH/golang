package main

import (
	"log"
	"net/http"
)

func handleRequests() {
	//http.HandleFunc("/", homePage)
	// add our articles route and map it to our
	// returnAllArticles function like so
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
