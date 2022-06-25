package shorturl

import (
	"regexp"
	"shorturl/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

var url1 = "google.com/1231231312"
var url3 = "https://yandex.ru/123"

func TestShorten(t *testing.T) {
	var ur Url
	case1, _, err := service.Shorturl.Shorten(&ur, url1)
	assert.NoError(t, err)
	assert.Regexp(t, regexp.MustCompile("^(http)(...)([a-z]+).(...).(.*)$"), case1)
}

func TestShortenWithScheme(t *testing.T) {
	var ur Url
	case3, _, err := service.Shorturl.Shorten(&ur, url3)
	assert.NoError(t, err)
	assert.Regexp(t, regexp.MustCompile("^(http)(...)([a-z]+).(...).(.*)$"), case3)
}
