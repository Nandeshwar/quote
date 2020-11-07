package api

import (
	"errors"
	"testing"

	"net/http"
	"net/http/httptest"

	"quote/pkg/model"
	"quote/pkg/service"
	mock_service "quote/pkg/service/mock"
	quoteService "quote/pkg/service/quote"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestApi(t *testing.T) {

	Convey("api test", t, func() {
		ctrl := gomock.NewController(t)
		eventDetailServiceMock := mock_service.NewMockIEventDetail(ctrl)
		infoServiceMock := mock_service.NewMockIInfo(ctrl)

		w := httptest.NewRecorder()
		s := NewServer(0, 0, ImageSize{}, "abc", 1, service.InfoEventService{}, quoteService.NewQuoteService())
		s.eventDetailService = eventDetailServiceMock
		s.infoService = infoServiceMock

		Convey(" /quotes-devotional ", func() {
			Convey("success:", func() {
				req := httptest.NewRequest("GET", "/quotes-devotional", nil)
				s.quotesDevotional(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, `<head><meta http-equiv='refresh' content='300' /> </head>`)
				So(w.Body.String(), ShouldContainSubstring, `<title>Quote</title>`)
				So(w.Body.String(), ShouldContainSubstring, `<a href='http://localhost:1922/image-motivational/mother-teresa-we-have-today.jpg' target='_blank'><img src='image-motivational/mother-teresa-we-have-today.jpg' alt='Nandeshwar' style='width:0px;height:0px;'> </a>`)
			})

			Convey("failure: POST is not allowed", func() {
				req := httptest.NewRequest("POST", "/quotes-devotional", nil)
				s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusMethodNotAllowed)
			})
		})

		Convey("/quotes-motivational ", func() {
			Convey("success:", func() {
				req := httptest.NewRequest("GET", "/quotes-motivational", nil)
				s.quotesMotivational(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, `<head>Quote for the day! <meta http-equiv='refresh' content='300' /> </head><h1>Quote for the day!</h1><title>Quote</title>`)
			})

			Convey("failure: POST is not allowed", func() {
				req := httptest.NewRequest("POST", "/quotes-motivational", nil)
				s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusMethodNotAllowed)
			})
		})

		Convey("api /events", func() {
			Convey("success: api /event/abc ", func() {

				eventDetailList := []model.EventDetail{{
					ID:    10,
					Day:   10,
					Month: 10,
					Year:  2020,
					Title: "title1",
				}}
				eventDetailServiceMock.EXPECT().GetEventDetailByTitleOrInfo("abc").Return(eventDetailList, nil)
				req := httptest.NewRequest("GET", "/events/abc", nil)
				s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "title1")
				So(w.Body.String(), ShouldContainSubstring, "10")
				So(w.Body.String(), ShouldContainSubstring, "Sat Oct 10, 2020")
			})

			Convey("success: api /event ", func() {

				eventDetailList := []model.EventDetail{{
					ID:    10,
					Day:   10,
					Month: 10,
					Year:  2020,
					Title: "title1",
				}}
				eventDetailServiceMock.EXPECT().GetEventDetailByTitleOrInfo("").Return(eventDetailList, nil)
				req := httptest.NewRequest("GET", "/events", nil)
				//s.events(w, req)
				s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "title1")
				So(w.Body.String(), ShouldContainSubstring, "10")
				So(w.Body.String(), ShouldContainSubstring, "Sat Oct 10, 2020")
			})

			Convey("failure: db error", func() {
				eventDetailServiceMock.EXPECT().GetEventDetailByTitleOrInfo("abc").Return(nil, errors.New("db error"))
				req := httptest.NewRequest("GET", "/events/abc", nil)
				s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})

			Convey("failure: POST is not allowed", func() {
				req := httptest.NewRequest("POST", "/events", nil)
				s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusMethodNotAllowed)
			})
		})

		Convey("api /info", func() {
			Convey("success: api /info/abc ", func() {

				infoList := []model.Info{{
					ID:    10,
					Title: "title1",
				}}
				infoServiceMock.EXPECT().GetInfoByTitleOrInfo("abc").Return(infoList, nil)
				req := httptest.NewRequest("GET", "/info/abc", nil)
				s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "title1")
				So(w.Body.String(), ShouldContainSubstring, "10")
			})

			Convey("failure: db error", func() {
				infoServiceMock.EXPECT().GetInfoByTitleOrInfo("abc").Return(nil, errors.New("db error"))
				req := httptest.NewRequest("GET", "/info/abc", nil)
				s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})

			Convey("failure: POST is not allowed", func() {
				req := httptest.NewRequest("POST", "/info", nil)
				s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusMethodNotAllowed)
			})
		})

		Convey("api /search", func() {
			Convey("success: api /search/abc ", func() {
				eventDetailList := []model.EventDetail{{
					ID:    10,
					Day:   10,
					Month: 10,
					Year:  2020,
					Title: "title1",
				}}

				infoList := []model.Info{{
					ID:    10,
					Title: "title1",
				}}
				infoServiceMock.EXPECT().GetInfoByTitleOrInfo("abc").Return(infoList, nil)
				eventDetailServiceMock.EXPECT().GetEventDetailByTitleOrInfo("abc").Return(eventDetailList, nil)

				req := httptest.NewRequest("GET", "/search/abc", nil)
				s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "title1")
				So(w.Body.String(), ShouldContainSubstring, "10")
			})

			Convey("failure: POST is not allowed", func() {
				req := httptest.NewRequest("POST", "/search/abc", nil)
				s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusMethodNotAllowed)
			})
		})

		Convey("getNonReadImageTest", func() {
			Convey("getUnique files in ", func() {

				allImagesPath := []string{
					"image/competitionWithMySelf.jpg",
					"image/becomegood.jpg",
					"image/thanks-to-obstacles.jpg",
					"image/you-are-responsible.jpg",
					"image/lift-up.jpg",
				}

				var imageRead []string

				for i := 0; i < len(allImagesPath); i++ {
					imageRead2, _ := getNonReadImage("test", len(allImagesPath), imageRead, s.quoteService.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, len(allImagesPath))

				for i := 0; i < len(allImagesPath); i++ {
					So(imageRead, ShouldContain, allImagesPath[i])
				}

				for i := 1; i <= len(allImagesPath)*2; i++ {
					imageRead2, _ := getNonReadImage("test", len(allImagesPath), imageRead, s.quoteService.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, len(allImagesPath))

				for i := 0; i < len(allImagesPath); i++ {
					So(imageRead, ShouldContain, allImagesPath[i])
				}

				imageRead = nil
				for i := 1; i <= len(allImagesPath)*2+1; i++ {
					imageRead2, _ := getNonReadImage("test", len(allImagesPath), imageRead, s.quoteService.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, 1)

				imageRead = nil
				for i := 1; i <= len(allImagesPath)*2+2; i++ {
					imageRead2, _ := getNonReadImage("test", len(allImagesPath), imageRead, s.quoteService.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, 2)

				imageRead = nil
				for i := 1; i <= len(allImagesPath)*2+3; i++ {
					imageRead2, _ := getNonReadImage("test", len(allImagesPath), imageRead, s.quoteService.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, 3)

				imageRead = nil
				for i := 1; i <= len(allImagesPath)*2+4; i++ {
					imageRead2, _ := getNonReadImage("test", len(allImagesPath), imageRead, s.quoteService.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, 4)

				imageRead = nil
				for i := 1; i <= len(allImagesPath)*2+5; i++ {
					imageRead2, _ := getNonReadImage("test", len(allImagesPath), imageRead, s.quoteService.GetQuoteMotivationalImage, allImagesPath)
					imageRead = imageRead2
				}
				So(len(imageRead), ShouldEqual, 5)

			})
		})

	})
}
