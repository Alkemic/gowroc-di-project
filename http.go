package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func listEntriesHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := fetchEntries()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println("got error fetching entries:", err)
		return
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(entries); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println("got error encoding entries:", err)
	}
}

func getEntryHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("id") == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("missing id parameter")
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("got error converting id to int:", err)
		return
	}

	entry, err := getEntry(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println("got error getting entry:", err)
		return
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(entry); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println("got error encoding entry:", err)
	}
}
