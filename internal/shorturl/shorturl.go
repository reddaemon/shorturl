package shorturl

import (
	"github.com/pkg/errors"
	"log"
	u "net/url"
	"strings"

	"shorturl/internal/utils"

	"github.com/speps/go-hashids"
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
	var sbUrl strings.Builder
	var sbShrlnk strings.Builder

	if !strings.Contains(url, "//") {
		sbUrl.WriteString("https://")
		sbUrl.WriteString(url)
	} else {
		sbUrl.WriteString(url)
	}

	temp, err := u.Parse(sbUrl.String())
	if err != nil {
		return "", sbUrl.String(), errors.Wrapf(err, "cannot parse url")
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

	shrtlnk, _ := hashId.Encode(hashSlice)
	sbShrlnk.WriteString(ur.Scheme)
	sbShrlnk.WriteString("://")
	sbShrlnk.WriteString(ur.Host)
	sbShrlnk.WriteString("/")
	sbShrlnk.WriteString(shrtlnk)

	log.Printf("Result string: %s", sbShrlnk.String())
	return sbShrlnk.String(), sbUrl.String(), err
}
