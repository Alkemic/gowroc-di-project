package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/posts", listEntriesHandler)
	http.HandleFunc("/post", getEntryHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
