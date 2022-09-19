package handlers

import (
	"log"
	"net/http"

	"shorturl/internal/service"
	"shorturl/internal/shorturl"
)

type handler struct {
	linkManager service.LinkManager
	shortener   shorturl.ShortenTool
}

type Handler interface {
	ShortHandler(w http.ResponseWriter, r *http.Request)
	GetFull(w http.ResponseWriter, r *http.Request)
}

func NewHandler(linkManager service.LinkManager,
	shortener shorturl.ShortenTool) Handler {
	return &handler{
		linkManager: linkManager,
		shortener:   shortener,
	}
}

func (h *handler) ShortHandler(w http.ResponseWriter, r *http.Request) {

	fullUrl := r.URL.Query().Get("url")

	shortUrl, fullUrl, err := h.shortener.Shorten(fullUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.linkManager.SetLink(shortUrl, fullUrl)
	if err != nil {
		log.Printf("Cannot save url to cache: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Cannot save short link: internal server error"))
		if err != nil {
			return
		}
		return
	}
	log.Printf("url saved, record id: %d", id)

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(shortUrl))
	if err != nil {
		return
	}
}

func (h *handler) GetFull(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Query().Get("shorturl")
	if shortUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("required parameter short url didn't specified"))
		if err != nil {
			return
		}
	}

	fullUrl, err := h.linkManager.GetLink(shortUrl)
	if err != nil {
		log.Printf("Cannot get full url by short url: %v", err)
		log.Printf("short, full: %v %v", shortUrl, fullUrl)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Cannot get full url by short url: internal server error"))
		if err != nil {
			return
		}
		return
	}

	http.Redirect(w, r, fullUrl, 301)
	w.WriteHeader(http.StatusOK)
}
