package service

import (
	"github.com/stretchr/testify/assert"
	"shorturl/internal/repository/mocks"
	"testing"
)

var setTests = []struct {
	short string
	full  string
}{
	{"http://localhost:8080/Jr", "google.com"},
	{"http://localhost:8080/Gk", "youtube.com"},
	{"http://localhost:8080/Sd", "mail.ru"},
}

var getTests = []struct {
	bucketName string
	short      string
	want       string
}{
	{"Urls", "http://localhost:8080/Jr", "https://youtube.com"},
	{"Urls", "http://localhost:8080/Gk", "https://google.com"},
	{"Urls", "http://localhost:8080/Sd", "https://mail.ru"},
}

func TestService_SetLink(t *testing.T) {
	mockRepo := new(mocks.RepoTool)
	service := service{mockRepo}

	for _, e := range setTests {
		mockRepo.On("Set", e.short, e.full).Return(int64(0), nil)
		_, err := service.SetLink(e.short, e.full)
		if err != nil {
			t.Fatal(err)
		}
		assert.NoError(t, err)

	}

}

func TestService_GetLink(t *testing.T) {
	mockRepo := new(mocks.RepoTool)
	service := service{mockRepo}

	for _, e := range getTests {
		mockRepo.On("Get", e.short).Return(e.want, nil)
		_, err := service.GetLink(e.short)
		if err != nil {
			t.Fatal(err)
		}
		assert.NoError(t, err)

	}

}
