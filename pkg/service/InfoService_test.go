package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"quote/pkg/model"
	mock_repo "quote/pkg/repo/mock"
)

func TestInfoService(t *testing.T) {
	Convey("Test Info Service", t, func() {
		ctrl := gomock.NewController(t)
		infoRepo := mock_repo.NewMockIInfoRepo(ctrl)
		quoteService := InfoEventService{
			InfoRepo: infoRepo,
		}

		Convey("Test ValidateForm", func() {
			quoteService := InfoEventService{}
			Convey("success: valid date 2020-10-16 13:32", func() {
				form := model.InfoForm{
					CreatedAt: "2020-10-16 13:32",
				}
				err := quoteService.ValidateForm(form)
				So(err, ShouldBeNil)
			})

			Convey("success: valid date time.Now()", func() {
				form := model.InfoForm{}
				err := quoteService.ValidateForm(form)
				So(err, ShouldBeNil)
			})

			Convey("failure: invalid date format- 2020", func() {
				form := model.InfoForm{
					CreatedAt: "2020",
				}
				err := quoteService.ValidateForm(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "wrong date and time format for createdAt. given date=2020, please provide date in this format yyyy-mm-dd tt:mm")
			})

			Convey("failure: invalid date format- 2020:10-16 13:30", func() {
				form := model.InfoForm{
					CreatedAt: "2020:10-16 13:30",
				}
				err := quoteService.ValidateForm(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "wrong date and time format for createdAt. given date=2020:10-16 13:30, please provide date in this format yyyy-mm-dd tt:mm")
			})

			Convey("failure: invalid date format- 2020-10:16 13:30", func() {
				form := model.InfoForm{
					CreatedAt: "2020-10:16 13:30",
				}
				err := quoteService.ValidateForm(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "wrong date and time format for createdAt. given date=2020-10:16 13:30, please provide date in this format yyyy-mm-dd tt:mm")
			})

			Convey("failure: invalid date format- 2020-10-16 13-30", func() {
				form := model.InfoForm{
					CreatedAt: "2020-10-16 13-30",
				}
				err := quoteService.ValidateForm(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "wrong date and time format for createdAt. given date=2020-10-16 13-30, please provide date in this format yyyy-mm-dd tt:mm")
			})

			Convey("failure: invalid date format- 9999-99-99 99:99", func() {
				form := model.InfoForm{
					CreatedAt: "9999-99-99 99:99",
				}
				err := quoteService.ValidateForm(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `parsing time "9999-99-99 99:99": month out of range`)
			})

			Convey("failure: invalid link- abc", func() {
				form := model.InfoForm{
					CreatedAt: "2020-10-16 13:30",
					Link:      "abc",
				}
				err := quoteService.ValidateForm(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "pipeline(|) separated links value must start with http or https. link could not be less than 4")
			})

			Convey("failure: invalid link- abcd", func() {
				form := model.InfoForm{
					CreatedAt: "2020-10-16 13:30",
					Link:      "abcd",
				}
				err := quoteService.ValidateForm(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "pipeline(|) separated links value must start with http or https")
			})

			Convey("failure: invalid link- http\"", func() {
				form := model.InfoForm{
					CreatedAt: "2020-10-16 13:30",
					Link:      "http\"",
				}
				err := quoteService.ValidateForm(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "pipeline(|) separated link's value should not ended with (\", ', .)")
			})

			Convey("failure: invalid link- http'", func() {
				form := model.InfoForm{
					CreatedAt: "2020-10-16 13:30",
					Link:      "http'",
				}
				err := quoteService.ValidateForm(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "pipeline(|) separated link's value should not ended with (\", ', .)")
			})

			Convey("failure: invalid link- http.", func() {
				form := model.InfoForm{
					CreatedAt: "2020-10-16 13:30",
					Link:      "http.",
				}
				err := quoteService.ValidateForm(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "pipeline(|) separated link's value should not ended with (\", ', .)")
			})
		})

		Convey("Test CreateNewInfo", func() {
			Convey("success: info creation", func() {
				form := model.InfoForm{
					Title:     "title1",
					CreatedAt: "2020-10-17 14:03",
					Link:      "http://google.com, http://yahoo.com",
				}

				infoRepo.EXPECT().CreateInfo(gomock.Any(), gomock.Any()).Return(int64(10), nil)
				id, err := quoteService.CreateNewInfo(context.Background(), form)
				So(err, ShouldBeNil)
				So(id, ShouldEqual, 10)
			})

			Convey("success: info creation with createdAt set with current date and time", func() {
				form := model.InfoForm{
					Title: "title1",
					Link:  "http://google.com, http://yahoo.com",
				}

				infoRepo.EXPECT().CreateInfo(context.Background(), gomock.Any()).Return(int64(10), nil)
				id, err := quoteService.CreateNewInfo(context.Background(), form)
				So(err, ShouldBeNil)
				So(id, ShouldEqual, 10)
			})

			Convey("failure: invalid date format 2020", func() {
				form := model.InfoForm{
					CreatedAt: "2020",
				}
				_, err := quoteService.CreateNewInfo(context.Background(), form)
				So(err, ShouldNotBeNil)
			})

			Convey("failure: info creation failed. db error", func() {
				form := model.InfoForm{
					Title:     "title1",
					CreatedAt: "2020-10-17 14:03",
					Link:      "http://google.com, http://yahoo.com",
				}

				infoRepo.EXPECT().CreateInfo(context.Background(), gomock.Any()).Return(int64(0), errors.New("db error"))
				id, err := quoteService.CreateNewInfo(context.Background(), form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "db error")
				So(id, ShouldEqual, 0)
			})
		})

		Convey("Test GetInfoByTitleOrInfo", func() {
			Convey("success: get info with 1 link", func() {
				infoList := []model.Info{{
					ID:    10,
					Title: "Title1",
					Info:  "Info1",
					Link:  "www.google.com",
				}}
				infoRepo.EXPECT().GetInfoByTitleOrInfo("abc").Return(infoList, nil)
				infoListFromDB, err := quoteService.GetInfoByTitleOrInfo("abc")

				So(err, ShouldBeNil)
				So(infoListFromDB[0].ID, ShouldEqual, infoList[0].ID)
				So(infoListFromDB[0].Title, ShouldEqual, infoList[0].Title)
				So(infoListFromDB[0].Info, ShouldEqual, infoList[0].Info)
				So(infoListFromDB[0].Links[0], ShouldEqual, infoList[0].Link)
			})

			Convey("success: get info with 2 link", func() {
				infoList := []model.Info{{
					ID:    10,
					Title: "Title1",
					Info:  "Info1",
					Link:  "www.google.com",
				}, {
					ID:    10,
					Title: "Title1",
					Info:  "Info1",
					Link:  "www.yahoo.com",
				}}
				infoRepo.EXPECT().GetInfoByTitleOrInfo("abc").Return(infoList, nil)
				infoListFromDB, err := quoteService.GetInfoByTitleOrInfo("abc")

				So(err, ShouldBeNil)
				So(len(infoListFromDB), ShouldEqual, 1)

				So(infoListFromDB[0].ID, ShouldEqual, infoList[0].ID)
				So(infoListFromDB[0].Title, ShouldEqual, infoList[0].Title)
				So(infoListFromDB[0].Info, ShouldEqual, infoList[0].Info)
				So(infoListFromDB[0].Links[0], ShouldEqual, infoList[0].Link)
				So(infoListFromDB[0].Links[1], ShouldEqual, infoList[1].Link)
			})

			Convey("success: get info with 3 link for id 10", func() {
				infoList := []model.Info{{
					ID:    10,
					Title: "Title1",
					Info:  "Info1",
					Link:  "www.google.com",
				}, {
					ID:    10,
					Title: "Title1",
					Info:  "Info1",
					Link:  "www.yahoo.com",
				}, {
					ID:    11,
					Title: "Title2",
					Info:  "Info2",
					Link:  "www.hotmail.com",
				}, {
					ID:    10,
					Title: "Title1",
					Info:  "Info1",
					Link:  "www.hotmail.com",
				}}
				infoRepo.EXPECT().GetInfoByTitleOrInfo("abc").Return(infoList, nil)
				infoListFromDB, err := quoteService.GetInfoByTitleOrInfo("abc")

				So(err, ShouldBeNil)
				So(len(infoListFromDB), ShouldEqual, 2)

				So(infoListFromDB[0].ID, ShouldEqual, infoList[0].ID)
				So(infoListFromDB[0].Title, ShouldEqual, infoList[0].Title)
				So(infoListFromDB[0].Info, ShouldEqual, infoList[0].Info)
				So(infoListFromDB[0].Links[0], ShouldEqual, infoList[0].Link)
				So(infoListFromDB[0].Links[1], ShouldEqual, infoList[1].Link)
				So(infoListFromDB[0].Links[2], ShouldEqual, infoList[2].Link)
			})

			Convey("failure: db error", func() {
				infoRepo.EXPECT().GetInfoByTitleOrInfo("abc").Return(nil, errors.New("db error"))
				_, err := quoteService.GetInfoByTitleOrInfo("abc")

				So(err, ShouldNotBeNil)

				So(err.Error(), ShouldEqual, "db error")
			})
		})
	})
}
