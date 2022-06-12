package handlers

import (
	"net/http"
	"shorturl/internal/shorturl"
	"shorturl/internal/usecase"
)

func ShortHandler(w http.ResponseWriter, r *http.Request) {
	var ur shorturl.Url

	u := r.URL.Query().Get("url")

	result, err := usecase.Shorturl.Shorten(&ur, u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
