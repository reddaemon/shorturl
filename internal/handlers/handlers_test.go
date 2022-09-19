package handlers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"shorturl/internal/service/linkManagerMocks"
	"shorturl/internal/shorturl"
	"testing"
)

var testUrls = []struct {
	url string
}{
	{"https://youtube.com"},
	{"https://google.com"},
	{"https://mail.ru"},
}

var shortUrls = []struct {
	url string
}{
	{"http://localhost:8080/GJ"},
	{"http://localhost:8080/vX"},
	{"http://localhost:8080/78"},
}

func TestShortHandler(t *testing.T) {
	linkManagerMock := linkManagerMocks.NewLinkManager(t)
	var url shorturl.Url

	w := httptest.NewRecorder()
	handlers := NewHandler(linkManagerMock, &url)
	for i, e := range testUrls {
		linkManagerMock.On("SetLink", shortUrls[i].url, e.url).Return(int64(0), nil)
		req := httptest.NewRequest(http.MethodPost,
			fmt.Sprintf("/v1/shorturl/short?url=%s", e.url), nil)
		handler := http.HandlerFunc(handlers.ShortHandler)
		handler.ServeHTTP(w, req)
		if status := w.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
		}
		fmt.Println(w.Body.String())

	}

}

func TestGetFull(t *testing.T) {
	linkManagerMock := linkManagerMocks.NewLinkManager(t)

	var url shorturl.Url
	handler := NewHandler(linkManagerMock, url)

	for i, e := range shortUrls {
		linkManagerMock.On("GetLink", e.url).Return(testUrls[i].url, nil)
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet,
			fmt.Sprintf("/v1/shorturl/full?shorturl=%s", e.url), nil)

		handler.GetFull(w, req)
		res := w.Result()
		resBody, _ := io.ReadAll(res.Body)

		assert.Equal(t, http.StatusMovedPermanently, w.Code)
		assert.Contains(t, string(resBody), "Moved Permanently")

	}
}
