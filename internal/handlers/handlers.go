package handlers

import (
	"log"
	"net/http"

	"shorturl/internal/service"
	"shorturl/internal/shorturl"
)
// сделать неэкспортируемым
type Handler struct {
	ServiceTool service.ServiceTool
	Shortener   *shorturl.Url
}
// изменить название
type HandlerTool interface {
	ShortHandler(w http.ResponseWriter, r *http.Request)
	GetFull(w http.ResponseWriter, r *http.Request)
}
// shortener должен быть интерфейсом
func NewHandler(serviceTool service.ServiceTool, shortener *shorturl.Url) *Handler {
	return &Handler{
		ServiceTool: serviceTool,
		Shortener:   shortener,
	}
}

func (h *Handler) ShortHandler(w http.ResponseWriter, r *http.Request) {

	fullUrl := r.URL.Query().Get("url")

	shortUrl, fullUrl, err := h.Shortener.Shorten(fullUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := h.ServiceTool.SetLink(shortUrl, fullUrl)
	if err != nil {
		log.Printf("Cannot save url to cache: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Cannot save shortlink: internal server error"))
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

func (h *Handler) GetFull(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Query().Get("shorturl")

	fullUrl, err := h.ServiceTool.GetLink(shortUrl)
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
