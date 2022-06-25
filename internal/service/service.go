package service

type Shorturl interface {
	Shorten(url string) (string, string, error)
}

type Repo interface {
	Set(short string, fullurl string) error
	Get(short string) (string, error)
	CreateBucket(name string) error
}
