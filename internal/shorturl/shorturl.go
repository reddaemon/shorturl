package shorturl

import (
	"github.com/pkg/errors"
	"log"
	u "net/url"
	"strings"

	"shorturl/internal/utils"

	"github.com/speps/go-hashids/v2"
)

type ShortenTool interface {
	Shorten(url string) (string, string, error)
}

type Url struct {
	Host   string
	Path   string
	Scheme string
}

func (ur Url) Shorten(url string) (string, string, error) {
	var builtUrl strings.Builder
	var builtShortLink strings.Builder

	if !strings.Contains(url, "//") {
		builtUrl.WriteString("https://")
		builtUrl.WriteString(url)
	} else {
		builtUrl.WriteString(url)
	}

	temp, err := u.Parse(builtUrl.String())
	if err != nil {
		return "", builtUrl.String(), errors.Wrapf(err, "cannot parse url")
	}

	log.Printf("temp: %s", temp)

	salt := utils.RandStringBytes(3)

	ur.Host = "localhost:8080"
	ur.Scheme = "http"

	hashSlice := []int{1}

	hashIdData := hashids.NewData()
	hashIdData.Salt = salt
	hashId, err := hashids.NewWithData(hashIdData)
	if err != nil {
		return "", "", errors.Wrapf(err, "cannot get hash")
	}

	shortLink, _ := hashId.Encode(hashSlice)
	builtShortLink.WriteString(ur.Scheme)
	builtShortLink.WriteString("://")
	builtShortLink.WriteString(ur.Host)
	builtShortLink.WriteString("/")
	builtShortLink.WriteString(shortLink)

	log.Printf("Result string: %s", builtShortLink.String())
	return builtShortLink.String(), builtUrl.String(), err
}
