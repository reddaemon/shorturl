package handlers

import (
	"log"
	"net/http"
	"shorturl/internal/repository"
	"shorturl/internal/service"
	"shorturl/internal/shorturl"
)

type Handler struct {
	repo repository.Repo
}

func NewHandler(repo *repository.Repo) *Handler {
	return &Handler{repo: *repo}
}

func (h *Handler) ShortHandler(w http.ResponseWriter, r *http.Request) {
	var ur shorturl.Url

	fullurl := r.URL.Query().Get("url")

	shortlink, fullurl, err := service.Shorturl.Shorten(&ur, fullurl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.repo.Set(shortlink, fullurl)
	if err != nil {
		log.Printf("Cannot get full url by short url: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cannot save shortlink: internal server error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(shortlink))
}

func (h *Handler) GetFull(w http.ResponseWriter, r *http.Request) {
	short := r.URL.Query().Get("shorturl")

	full, err := h.repo.Get("Urls", short)
	if err != nil {
		log.Printf("Cannot get full url by short url: %v", err)
		log.Printf("short, full: %v %v", short, full)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cannot get full url by short url: internal server error"))
		return
	}

	http.Redirect(w, r, string(full), 301)
	w.WriteHeader(http.StatusOK)
}
