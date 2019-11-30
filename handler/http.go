package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Alkemic/gowroc-di-project/repository"
)

type blogService interface {
	List() ([]repository.Post, error)
	Get(id int) (repository.Post, error)
}

type httpHandler struct {
	blogService blogService
}

func NewHandler(blogService blogService) http.Handler {
	h := &httpHandler{blogService: blogService}
	mux := http.NewServeMux()
	mux.HandleFunc("/posts", h.listEntriesHandler)
	mux.HandleFunc("/post", h.getEntryHandler)
	return mux
}

func (h httpHandler) listEntriesHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := h.blogService.List()
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

func (h httpHandler) getEntryHandler(w http.ResponseWriter, r *http.Request) {
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

	entry, err := h.blogService.Get(id)
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
