package service

import (
	"errors"
	mock_repo "quote/pkg/repo/mock"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	sqliteRepo := mock_repo.NewMockIRepo(mockCtrl)
	q := QuoteService{SQLite3Repo: sqliteRepo}

	Convey("Test login", t, func() {

		Convey("success: login", func() {
			sqliteRepo.EXPECT().LoginInfo("abc", "xyz").Return(nil)
			err := q.Login("abc", "xyz")
			So(err, ShouldBeNil)
		})

		Convey("fail: login", func() {
			sqliteRepo.EXPECT().LoginInfo("abc", "xyz").Return(errors.New("invalid user name and password"))
			err := q.Login("abc", "xyz")
			So(err.Error(), ShouldEqual, "invalid user name and password")
		})

	})
}
