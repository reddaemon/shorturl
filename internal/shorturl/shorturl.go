package shorturl

import (
	"fmt"
	"log"
	u "net/url"
	"strings"

	"github.com/speps/go-hashids"
)

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
		return "", sbUrl.String(), err
	}

	log.Printf("temp: %s", temp)

	/*if temp.Path == "" {
		return "", sbUrl.String(), errors.New("path is empty, nothing to do...please add path")
	}*/

	ur.Host = "localhost:8080"
	ur.Path = temp.Path
	ur.Scheme = "http"

	hashSlice := []int{1}

	hd := hashids.NewData()
	hd.Salt = ur.Path
	h, err := hashids.NewWithData(hd)
	if err != nil {
		fmt.Println(err)
	}

	shrtlnk, _ := h.Encode(hashSlice)
	sbShrlnk.WriteString(ur.Scheme)
	sbShrlnk.WriteString("://")
	sbShrlnk.WriteString(ur.Host)
	sbShrlnk.WriteString("/")
	sbShrlnk.WriteString(shrtlnk)

	log.Printf("Result string: %s", sbShrlnk.String())
	return sbShrlnk.String(), sbUrl.String(), err
}
