package service

import (
	"github.com/pkg/errors"
	"shorturl/internal/repository"
)

// https://dave.cheney.net/practical-go/presentations/qcon-china.html

// todo изменить название
type ServiceTool interface {
	SetLink(shortlink string, fullurl string) (id int64, err error)
	GetLink(shorturl string) (string, error)
}

type service struct {
	repoTool repository.RepoTool
}

func NewService(repoTool repository.RepoTool) ServiceTool {
	return &service{repoTool: repoTool}
}

func (s *service) SetLink(shortUrl string, fullUrl string) (id int64, err error) {
	id, err = s.repoTool.Set(shortUrl, fullUrl)
	if err != nil {
		return id, errors.Wrapf(err, "cannot save shortUrl, short and full: %s %s", shortUrl, fullUrl)
	}
	return id, nil
}

func (s *service) GetLink(shortUrl string) (fullUrl string, err error) {
	fullUrl, err = s.repoTool.Get(shortUrl)
	if err != nil {
		return fullUrl, errors.Wrapf(err, "cannot get fullUrl by shortUrl %s", shortUrl)
	}
	return fullUrl, nil
}
