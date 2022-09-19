package service

import (
	"github.com/pkg/errors"
	"shorturl/internal/repository"
)

// https://dave.cheney.net/practical-go/presentations/qcon-china.html

type LinkManager interface {
	SetLink(shortLink string, fullUrl string) (id int64, err error)
	GetLink(shortUrl string) (string, error)
}

type service struct {
	repository repository.Repository
}

func NewLinkManager(repoTool repository.Repository) LinkManager {
	return &service{repository: repoTool}
}

func (s *service) SetLink(shortUrl string, fullUrl string) (id int64, err error) {
	id, err = s.repository.Set(shortUrl, fullUrl)
	if err != nil {
		return id, errors.Wrapf(err, "cannot save shortUrl, short and full: %s %s", shortUrl, fullUrl)
	}
	return id, nil
}

func (s *service) GetLink(shortUrl string) (fullUrl string, err error) {
	fullUrl, err = s.repository.Get(shortUrl)
	if err != nil {
		return fullUrl, errors.Wrapf(err, "cannot get fullUrl by shortUrl %s", shortUrl)
	}
	return fullUrl, nil
}
