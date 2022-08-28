package shorturl

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"shorturl/internal/mock_shorturl"
	"testing"

	"github.com/golang/mock/gomock"
)

var urlTests = []struct {
	url string
}{
	{"google.com/1231231312"},
	{"https://yandex.ru/123"},
}

func TestShorten(t *testing.T) {
	ctrl := gomock.NewController(t)
	sTool := mock_shorturl.NewMockShortenTool(ctrl)
	var url Url
	for _, e := range urlTests {
		sTool.EXPECT().
			Shorten(e.url).
			Return(gomock.Any().String(), gomock.Any().String(), nil).
			AnyTimes()
		shortUrl, _, err := url.Shorten(e.url)
		assert.NoError(t, err)
		assert.Regexp(t, regexp.MustCompile("^(http)(...)([a-z]+).(...).(.*)$"), shortUrl)
	}
}
