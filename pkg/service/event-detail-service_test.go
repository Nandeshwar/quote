package service

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"

	"quote/pkg/model"
	mock_repo "quote/pkg/repo/mock"
)

func TestEventDetailSerivce(t *testing.T) {
	Convey("Test Event Detail Service", t, func() {
		ctrl := gomock.NewController(t)
		eventInfoRepo := mock_repo.NewMockIEventDetailRepo(ctrl)

		quoteService := InfoEventService{
			EventDetailRepo: eventInfoRepo,
		}

		Convey("Test ValidateFormEvent", func() {
			quoteService := InfoEventService{}
			Convey("success: validate form", func() {
				form := model.EventDetailForm{
					EventDate: "2020-10-18",
					Typ:       "different",
					CreatedAt: "2010-10-18 17:10",
				}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldBeNil)
			})

			Convey("failure: invalid date - 2020-10. day is not provided", func() {
				form := model.EventDetailForm{
					EventDate: "2020-10",
				}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, fmt.Sprintf("wrong date and time format for event date. given date=%s, please provide date in this format yyyy-mm-dd", form.EventDate))
			})

			Convey("failure: invalid date - 2020/10-18. delimeter should be hyphen", func() {
				form := model.EventDetailForm{
					EventDate: "2020/10-18",
				}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, fmt.Sprintf("wrong date and time format for event date. given date=%s, please provide date in this format yyyy-mm-dd", form.EventDate))
			})

			Convey("failure: invalid date - 2020-10/18. delimeter should be hyphen", func() {
				form := model.EventDetailForm{
					EventDate: "20200-10/18",
				}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, fmt.Sprintf("wrong date and time format for event date. given date=%s, please provide date in this format yyyy-mm-dd", form.EventDate))
			})

			Convey("failure: empty event date is not allowed", func() {
				form := model.EventDetailForm{}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "event date should not be empty")
			})

			Convey("failure: invalid event date - 9999-99-99", func() {
				form := model.EventDetailForm{
					EventDate: "9999-99-99",
				}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `parsing time "9999-99-99": month out of range`)
			})

			Convey("failure: wrong event detail type- same1", func() {
				form := model.EventDetailForm{
					EventDate: "2020-10-18",
					Typ:       "same1",
				}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, fmt.Sprintf("invalid value for event type. given type=%s expected value: same/different", form.Typ))
			})

			Convey("failure: wrong date format for createdAt- 2020", func() {
				form := model.EventDetailForm{
					EventDate: "2020-10-18",
					Typ:       "same",
					CreatedAt: "2010",
				}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, fmt.Sprintf("wrong date and time format for createdAt. given date=%s, please provide date in this format yyyy-mm-dd tt:mm", form.CreatedAt))
			})

			Convey("failure: wrong date format for createdAt- 2010/10-18 17:10", func() {
				form := model.EventDetailForm{
					EventDate: "2020-10-18",
					Typ:       "same",
					CreatedAt: "2010/10-18 17:10",
				}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, fmt.Sprintf("wrong date and time format for createdAt. given date=%s, please provide date in this format yyyy-mm-dd tt:mm", form.CreatedAt))
			})

			Convey("failure: wrong date format for createdAt- 2010-10/18 17:10", func() {
				form := model.EventDetailForm{
					EventDate: "2020-10-18",
					Typ:       "same",
					CreatedAt: "2010-10/18 17:10",
				}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, fmt.Sprintf("wrong date and time format for createdAt. given date=%s, please provide date in this format yyyy-mm-dd tt:mm", form.CreatedAt))
			})

			Convey("failure: wrong date format for createdAt- 2010-10-18 17/10", func() {
				form := model.EventDetailForm{
					EventDate: "2020-10-18",
					Typ:       "same",
					CreatedAt: "2010-10-18 17/10",
				}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, fmt.Sprintf("wrong date and time format for createdAt. given date=%s, please provide date in this format yyyy-mm-dd tt:mm", form.CreatedAt))
			})
		})

		Convey("test validateLink", func() {
			Convey("failure: invalid link- abcd", func() {
				form := model.EventDetailForm{
					CreatedAt: "2020-10-16 10:10",
					EventDate: "2020-10-10",
					Link:      "abcd",
				}
				err := quoteService.ValidateFormEvent(form)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "pipeline(|) separated links value must start with http or https")
			})
		})

		Convey("Test CreateNewEventDetail", func() {

			form := model.EventDetailForm{
				EventDate: "2020-10-18",
				Title:     "Title1",
				Info:      "Event Detail",
				Typ:       "same",
				Link:      "http://google.com|http://yahoo.com",
				CreatedAt: "2020-10-19 10:19",
			}

			Convey("success: create new event detail", func() {
				eventInfoRepo.EXPECT().CreateEventDetail(gomock.Any()).Return(int64(10), nil)
				id, err := quoteService.CreateNewEventDetail(form)
				So(err, ShouldBeNil)
				So(id, ShouldEqual, 10)
			})

			Convey("success: create new event detail with current date and time for createdAt", func() {
				form2 := model.EventDetailForm{
					EventDate: "2020-10-18",
					Title:     "Title1",
					Info:      "Event Detail",
					Typ:       "same",
				}

				eventInfoRepo.EXPECT().CreateEventDetail(gomock.Any()).Return(int64(10), nil)
				id, err := quoteService.CreateNewEventDetail(form2)
				So(err, ShouldBeNil)
				So(id, ShouldEqual, 10)
			})

			Convey("success: create new event detail with default event type", func() {
				form2 := model.EventDetailForm{
					EventDate: "2020-10-18",
					Title:     "Title1",
					Info:      "Event Detail",
				}

				eventInfoRepo.EXPECT().CreateEventDetail(gomock.Any()).Return(int64(10), nil)
				id, err := quoteService.CreateNewEventDetail(form2)
				So(err, ShouldBeNil)
				So(id, ShouldEqual, 10)
			})

			Convey("failure: create new event detail", func() {
				eventInfoRepo.EXPECT().CreateEventDetail(gomock.Any()).Return(int64(0), errors.New("db error"))
				id, err := quoteService.CreateNewEventDetail(form)
				So(err, ShouldNotBeNil)
				So(id, ShouldEqual, 0)
				So(err.Error(), ShouldEqual, "db error")
			})

			Convey("failure: wrong date format for createdAt", func() {
				form := model.EventDetailForm{
					EventDate: "2020-10-18",
					Title:     "Title1",
					Info:      "Event Detail",
					Typ:       "same",
					CreatedAt: "2020-10-19 1019",
				}
				id, err := quoteService.CreateNewEventDetail(form)
				So(err, ShouldNotBeNil)
				So(id, ShouldEqual, 0)
				So(err.Error(), ShouldEqual, `parsing time "2020-10-19 1019" as "2006-01-02 15:04": cannot parse "19" as ":"`)
			})

			Convey("failure: wrong date format for event date", func() {
				form := model.EventDetailForm{
					EventDate: "2020-10:18",
					Title:     "Title1",
					Info:      "Event Detail",
					Typ:       "same",
					CreatedAt: "2020-10-19 1019",
				}
				id, err := quoteService.CreateNewEventDetail(form)
				So(err, ShouldNotBeNil)
				So(id, ShouldEqual, 0)
				So(err.Error(), ShouldEqual, `parsing time "2020-10:18" as "2006-01-02": cannot parse ":18" as "-"`)
			})
		})

		Convey("Test GetEventDetailByTitleOrInfo", func() {
			Convey("success: get 1 event 1 link", func() {
				eventDetailFromDb := []model.EventDetail{{
					ID:    10,
					Title: "Title1",
					URL:   "www.google.com",
				}}
				eventInfoRepo.EXPECT().GetEventDetailByTitleOrInfo("abc").Return(eventDetailFromDb, nil)
				eventDetailList, err := quoteService.GetEventDetailByTitleOrInfo("abc")
				So(err, ShouldBeNil)
				So(len(eventDetailList), ShouldEqual, 1)
				So(eventDetailList[0].ID, ShouldEqual, 10)
				So(len(eventDetailList[0].Links), ShouldEqual, 1)
				So(eventDetailList[0].Links[0], ShouldEqual, "www.google.com")
			})

			Convey("success: get 1 event 3 links", func() {
				eventDetailFromDb := []model.EventDetail{{
					ID:    10,
					Title: "Title1",
					URL:   "www.google.com",
				}, {
					ID:    10,
					Title: "Title1",
					URL:   "www.yahoo.com",
				}, {
					ID:    10,
					Title: "Title1",
					URL:   "www.hotmail.com",
				}}
				eventInfoRepo.EXPECT().GetEventDetailByTitleOrInfo("abc").Return(eventDetailFromDb, nil)
				eventDetailList, err := quoteService.GetEventDetailByTitleOrInfo("abc")
				So(err, ShouldBeNil)
				So(len(eventDetailList), ShouldEqual, 1)
				So(eventDetailList[0].ID, ShouldEqual, 10)
				So(len(eventDetailList[0].Links), ShouldEqual, 3)
				So(eventDetailList[0].Links[0], ShouldEqual, "www.google.com")
				So(eventDetailList[0].Links[1], ShouldEqual, "www.yahoo.com")
				So(eventDetailList[0].Links[2], ShouldEqual, "www.hotmail.com")
			})

			Convey("success: get 2 events 3,1 links", func() {
				eventDetailFromDb := []model.EventDetail{{
					ID:    10,
					Title: "Title1",
					URL:   "www.google.com",
				}, {
					ID:    10,
					Title: "Title1",
					URL:   "www.yahoo.com",
				}, {
					ID:    11,
					Title: "Title1",
					URL:   "www.yahoo.com",
				}, {
					ID:    10,
					Title: "Title1",
					URL:   "www.hotmail.com",
				}}
				eventInfoRepo.EXPECT().GetEventDetailByTitleOrInfo("abc").Return(eventDetailFromDb, nil)
				eventDetailList, err := quoteService.GetEventDetailByTitleOrInfo("abc")
				So(err, ShouldBeNil)
				So(len(eventDetailList), ShouldEqual, 2)
				So(eventDetailList[0].ID, ShouldEqual, 10)
				So(len(eventDetailList[0].Links), ShouldEqual, 3)
				So(eventDetailList[0].Links[0], ShouldEqual, "www.google.com")
				So(eventDetailList[0].Links[1], ShouldEqual, "www.yahoo.com")
				So(eventDetailList[0].Links[2], ShouldEqual, "www.hotmail.com")
				So(eventDetailList[1].Links[0], ShouldEqual, "www.yahoo.com")
			})

			Convey("failure: db error", func() {
				eventInfoRepo.EXPECT().GetEventDetailByTitleOrInfo("abc").Return(nil, errors.New("db error"))
				_, err := quoteService.GetEventDetailByTitleOrInfo("abc")
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "db error")
			})
		})

		Convey("Test GetEventDetailByMonthDay", func() {
			Convey("success: get 1 event 1 link", func() {
				eventDetailFromDb := []model.EventDetail{{
					ID:    10,
					Title: "Title1",
					URL:   "www.google.com",
				}}

				eventInfoRepo.EXPECT().GetEventDetailByMonthDay(2020, 10).Return(eventDetailFromDb, nil)
				eventDetailList, err := quoteService.GetEventDetailByMonthDay(2020, 10)

				So(err, ShouldBeNil)
				So(len(eventDetailList), ShouldEqual, 1)
				So(eventDetailList[0].ID, ShouldEqual, 10)
				So(len(eventDetailList[0].Links), ShouldEqual, 1)
				So(eventDetailList[0].Links[0], ShouldEqual, "www.google.com")
			})

			Convey("success: get 1 event 3 links", func() {
				eventDetailFromDb := []model.EventDetail{{
					ID:    10,
					Title: "Title1",
					URL:   "www.google.com",
				}, {
					ID:    10,
					Title: "Title1",
					URL:   "www.yahoo.com",
				}, {
					ID:    10,
					Title: "Title1",
					URL:   "www.hotmail.com",
				}}
				eventInfoRepo.EXPECT().GetEventDetailByMonthDay(2020, 10).Return(eventDetailFromDb, nil)
				eventDetailList, err := quoteService.GetEventDetailByMonthDay(2020, 10)
				So(err, ShouldBeNil)
				So(len(eventDetailList), ShouldEqual, 1)
				So(eventDetailList[0].ID, ShouldEqual, 10)
				So(len(eventDetailList[0].Links), ShouldEqual, 3)
				So(eventDetailList[0].Links[0], ShouldEqual, "www.google.com")
				So(eventDetailList[0].Links[1], ShouldEqual, "www.yahoo.com")
				So(eventDetailList[0].Links[2], ShouldEqual, "www.hotmail.com")
			})

			Convey("success: get 2 events 3,1 links", func() {
				eventDetailFromDb := []model.EventDetail{{
					ID:    10,
					Title: "Title1",
					URL:   "www.google.com",
				}, {
					ID:    10,
					Title: "Title1",
					URL:   "www.yahoo.com",
				}, {
					ID:    11,
					Title: "Title1",
					URL:   "www.yahoo.com",
				}, {
					ID:    10,
					Title: "Title1",
					URL:   "www.hotmail.com",
				}}
				eventInfoRepo.EXPECT().GetEventDetailByMonthDay(2020, 10).Return(eventDetailFromDb, nil)
				eventDetailList, err := quoteService.GetEventDetailByMonthDay(2020, 10)
				So(err, ShouldBeNil)
				So(len(eventDetailList), ShouldEqual, 2)
				So(eventDetailList[0].ID, ShouldEqual, 10)
				So(len(eventDetailList[0].Links), ShouldEqual, 3)
				So(eventDetailList[0].Links[0], ShouldEqual, "www.google.com")
				So(eventDetailList[0].Links[1], ShouldEqual, "www.yahoo.com")
				So(eventDetailList[0].Links[2], ShouldEqual, "www.hotmail.com")
				So(eventDetailList[1].Links[0], ShouldEqual, "www.yahoo.com")
			})

			Convey("failure: db error", func() {
				eventInfoRepo.EXPECT().GetEventDetailByMonthDay(2020, 10).Return(nil, errors.New("db error"))
				_, err := quoteService.GetEventDetailByMonthDay(2020, 10)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "db error")
			})
		})
		Convey("Test EventsInFuture", func() {
			Convey("success: events in future", func() {

				t := time.Now()
				year, month, day := t.Date()

				eventDetailFromDb := []model.EventDetail{{
					ID:    10,
					Title: "Title1",
					URL:   "www.google.com",
					Month: int(month),
					Day:   day,
					Year:  year,
				}}

				eventInfoRepo.EXPECT().GetEventDetailByMonthDay(int(month), day).Return(eventDetailFromDb, nil)

				todayEvents, err := quoteService.EventsInFuture(t)
				So(err, ShouldBeNil)
				So(len(todayEvents), ShouldEqual, 1)
				So(todayEvents[0].ID, ShouldEqual, 10)
				So(todayEvents[0].Links[0], ShouldEqual, "www.google.com")
				So(todayEvents[0].Year, ShouldEqual, year)
				So(todayEvents[0].Month, ShouldEqual, int(month))
			})

			Convey("success: events in future 0 records year does not match for type different", func() {

				t := time.Now()
				year, month, day := t.Date()

				eventDetailFromDb := []model.EventDetail{{
					ID:    10,
					Title: "Title1",
					URL:   "www.google.com",
					Month: int(month),
					Type:  "different",
					Day:   day,
					Year:  year - 1,
				}}

				eventInfoRepo.EXPECT().GetEventDetailByMonthDay(int(month), day).Return(eventDetailFromDb, nil)

				todayEvents, err := quoteService.EventsInFuture(t)
				So(err, ShouldBeNil)
				So(len(todayEvents), ShouldEqual, 0)
			})

			Convey("failure: db error", func() {
				t := time.Now()
				_, month, day := t.Date()
				eventInfoRepo.EXPECT().GetEventDetailByMonthDay(int(month), day).Return(nil, errors.New("db error"))

				_, err := quoteService.EventsInFuture(t)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "db error")
			})
		})
	})
}
