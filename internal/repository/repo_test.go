package repository

import (
	"shorturl/internal/mock_service"
	"testing"

	"github.com/golang/mock/gomock"
	//"github.com/stretchr/testify/assert"
)

var setTests = []struct {
	short string
	full string
}{
	{"http://localhost:8080/Jr","google.com"},
	{"http://localhost:8080/Gk","youtube.com"},
	{"http://localhost:8080/Sd","mail.ru"},
}

var getTests = []struct {
	bucketName string
	short string
	want string
}{
	{"Urls","http://localhost:8080/Jr", "https://youtube.com"},
	{"Urls","http://localhost:8080/Gk", "https://google.com"},
	{"Urls","http://localhost:8080/Sd", "https://mail.ru"},
}

func TestSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mRepo := mock_service.NewMockRepo(ctrl)
	for _, e := range setTests {
		mRepo.EXPECT().
		Set(e.short, e.full).
		Return(nil).
		AnyTimes()
	}
	
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mRepo := mock_service.NewMockRepo(ctrl)
	for _, e := range getTests {
		mRepo.EXPECT().Get(gomock.Eq(e.bucketName), gomock.Eq(e.short)).
		Return([]byte(e.want), nil).
		AnyTimes()
	}
	
}