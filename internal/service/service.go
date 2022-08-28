package service

import (
	"github.com/pkg/errors"
	"shorturl/internal/repository"
)

type ServiceTool interface {
	SetLink(shortlink string, fullurl string) (id int64, err error)
	GetLink(shorturl string) (string, error)
}

type Service struct {
	RepoTool repository.RepoTool
}

func NewService(repoTool repository.RepoTool) *Service {
	return &Service{
		RepoTool: repoTool}
}

func (s *Service) SetLink(shortUrl string, fullUrl string) (id int64, err error) {
	id, err = s.RepoTool.Set(shortUrl, fullUrl)
	if err != nil {
		return id, errors.Wrapf(err, "cannot save shortUrl, short and full: %s %s", shortUrl, fullUrl)
	}
	return id, nil
}

func (s *Service) GetLink(shortUrl string) (fullUrl string, err error) {
	fullUrl, err = s.RepoTool.Get(shortUrl)
	if err != nil {
		return fullUrl, errors.Wrapf(err, "cannot get fullUrl by shortUrl %s", shortUrl)
	}
	return fullUrl, nil
}
