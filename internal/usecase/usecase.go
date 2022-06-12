package usecase

type Shorturl interface {
	Shorten(url string) (string, error)
}
