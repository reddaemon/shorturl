package shorturl

import (
	"regexp"
	"shorturl/internal/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

var url1 = "google.com/1231231312"
var url2 = "google.com"
var url3 = "https://yandex.ru/123"

func TestShorten(t *testing.T) {
	var ur Url
	case1, err := usecase.Shorturl.Shorten(&ur, url1)
	assert.NoError(t, err)
	assert.Regexp(t, regexp.MustCompile("^(https)(...)([a-z]+).(...).(.*)$"), case1)
}

func TestShortenEmptyPathError(t *testing.T) {
	var ur Url
	_, err := usecase.Shorturl.Shorten(&ur, url2)
	assert.ErrorContains(t, err, "path is empty, nothing to do...please add path")
}

func TestShortenWithScheme(t *testing.T) {
	var ur Url
	case3, err := usecase.Shorturl.Shorten(&ur, url3)
	assert.NoError(t, err)
	assert.Regexp(t, regexp.MustCompile("^(https)(...)([a-z]+).(...).(.*)$"), case3)
}
