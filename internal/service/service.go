package service

type Shorturl interface {
	Shorten(url string) (string, string, error)
}

type Repo interface {
	Set(short string, fullurl string) error
	Get(BucketName string, short string) (fullurl []byte, err error)
	CreateBucket(BucketName string) error
}
