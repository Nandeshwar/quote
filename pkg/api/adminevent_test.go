package api

import (
	"errors"
	"github.com/gorilla/sessions"
	"net/http"
	"net/http/httptest"
	"quote/pkg/model"
	"strings"
	"testing"

	"quote/pkg/service"
	mock_service "quote/pkg/service/mock"
	"quote/pkg/service/quote"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAdminEvent(t *testing.T) {
	Convey("Test Admin Event API", t, func() {
		w := httptest.NewRecorder()
		s := NewServer(0, ImageSize{}, "abc", 1, service.InfoEventService{}, quote.NewQuoteService())
		Convey("GET API", func() {
			Convey("success: Get /admin-event-detail", func() {
				s.views.AdminEventDetail = "../../views/admin-event-detail.gtpl"
				req := httptest.NewRequest("GET", "/admin-event-detail", nil)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminEvent(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "Title")
				So(w.Body.String(), ShouldContainSubstring, "Info")
				So(w.Body.String(), ShouldContainSubstring, "Links Pipeline(|) Separated")
				So(w.Body.String(), ShouldContainSubstring, "Created Date")
			})

			Convey("failure: Get /admin-event-detail view not found", func() {
				s.views.AdminInfo = "/login1.gtpl"
				req := httptest.NewRequest("GET", "/admin-event-detail", nil)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminEvent(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusNotFound)
			})

			Convey("failure: Get /admin-event-detail authentication false", func() {
				s.views.AdminInfo = "../../views/admin-event-detail.gtpl"
				req := httptest.NewRequest("GET", "/admin-event-detail", nil)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = false
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminEvent(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusForbidden)
			})

			Convey("failure: Get /admin-event-detail wrong cookie store ", func() {
				s.views.AdminInfo = "../../views/admin-event-detail.gtpl"
				req := httptest.NewRequest("GET", "/admin-event-detail", nil)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie2")

				session.Values["authenticated"] = false
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminEvent(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusForbidden)
			})

			Convey("failure: Get /admin-event-detail invalid session ", func() {
				s.views.AdminInfo = "../../views/admin-event-detail.gtpl"
				req := httptest.NewRequest("GET", "/admin-event-detail", nil)

				s.sessionCookieStore = sessions.NewCookieStore([]byte("nandeshwar"))
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie2")

				session.Values["authenticated"] = false
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminEvent(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusForbidden)
			})

			Convey("failure: Get /admin-event-detail empty cookie name", func() {
				s.views.AdminEventDetail = "../../views/admin-event-detail.gtpl"
				req := httptest.NewRequest("GET", "/admin-event-detail", nil)

				s.cookieName = ""

				s.adminEvent(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusForbidden)
			})

			Convey("failure: Get /admin-info invalid html", func() {
				s.views.AdminEventDetail = "testdata/invalid-html-file.gtpl"
				req := httptest.NewRequest("GET", "/admin-event-detail", nil)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminEvent(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})
		})
		Convey("Post API", func() {
			ctrl := gomock.NewController(t)
			eventDetailService := mock_service.NewMockIEventDetail(ctrl)
			s.eventDetailService = eventDetailService
			Convey("success: post /admin-event-detail", func() {
				s.views.AdminEventDetail = "../../views/admin-event-detail.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=http://www.google.com&eventDate=2020-10-25&createdAt=2020-10-27 08:08&eventType=same")
				req := httptest.NewRequest("POST", "/admin-event-detail", body)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				infoForm := model.EventDetailForm{
					Title:     "title1",
					Info:      "info1",
					EventDate: "2020-10-25",
					Link:      "http://www.google.com",
					Typ:       "same",
					CreatedAt: "2020-10-27 08:08",
				}
				eventDetailService.EXPECT().ValidateFormEvent(infoForm).Return(nil)
				eventDetailService.EXPECT().CreateNewEventDetail(infoForm).Return(int64(10), nil)

				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				s.adminEvent(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "10")
			})

			Convey("failure: post /admin-event-detail missing field createdAt", func() {
				s.views.AdminEventDetail = "../../views/admin-event-detail.gtpl"
				body := strings.NewReader("title=title1&info=info&link=http://www.google.com&createdAt1=2020-10-25 11:42")
				req := httptest.NewRequest("POST", "/admin-event-detail", body)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminEvent(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusNotFound)
			})

			Convey("failure: post /admin-event-detail invalid form", func() {
				s.views.AdminEventDetail = "../../views/admin-event-detail.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=abc&createdAt=2020-10-25 11:42&eventDate=2010-10-25&eventType=same")
				req := httptest.NewRequest("POST", "/admin-event-detail", body)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				eventDetailForm := model.EventDetailForm{
					Title:     "title1",
					Info:      "info1",
					Link:      "abc",
					CreatedAt: "2020-10-25 11:42",
					EventDate: "2010-10-25",
					Typ:       "same",
				}
				eventDetailService.EXPECT().ValidateFormEvent(eventDetailForm).Return(errors.New("invalid link"))

				s.adminEvent(w, req)
				resp := w.Result()

				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "validation error. check log")
			})

			Convey("failure: post /admin-event-detail db error creating record", func() {
				s.views.AdminEventDetail = "../../views/admin-event-detail.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=http://www.google.com&createdAt=2020-10-25 11:42&eventDate=2010-10-25&eventType=same")
				req := httptest.NewRequest("POST", "/admin-event-detail", body)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				eventDetailForm := model.EventDetailForm{
					Title:     "title1",
					Info:      "info1",
					Link:      "http://www.google.com",
					CreatedAt: "2020-10-25 11:42",
					EventDate: "2010-10-25",
					Typ:       "same",
				}
				eventDetailService.EXPECT().ValidateFormEvent(eventDetailForm).Return(nil)
				eventDetailService.EXPECT().CreateNewEventDetail(eventDetailForm).Return(int64(0), errors.New("db error"))
				s.adminEvent(w, req)
				resp := w.Result()

				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "Did not create record. error. check log")
			})

			Convey("failure: post /admin-event-detail invalid html", func() {
				s.views.AdminEventDetail = "testdata/invalid-admin-event-detail.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=http://www.google.com&eventDate=2020-10-25&createdAt=2020-10-27 08:08&eventType=same")
				req := httptest.NewRequest("POST", "/admin-event-detail", body)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				infoForm := model.EventDetailForm{
					Title:     "title1",
					Info:      "info1",
					EventDate: "2020-10-25",
					Link:      "http://www.google.com",
					Typ:       "same",
					CreatedAt: "2020-10-27 08:08",
				}
				eventDetailService.EXPECT().ValidateFormEvent(infoForm).Return(nil)
				eventDetailService.EXPECT().CreateNewEventDetail(infoForm).Return(int64(10), nil)

				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				s.adminEvent(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})

			Convey("failure: post /admin-event-detail invalid form - invalid html", func() {
				s.views.AdminEventDetail = "testdata/invalid-admin-event-detail.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=abc&createdAt=2020-10-25 11:42&eventDate=2010-10-25&eventType=same")
				req := httptest.NewRequest("POST", "/admin-event-detail", body)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				eventDetailForm := model.EventDetailForm{
					Title:     "title1",
					Info:      "info1",
					Link:      "abc",
					CreatedAt: "2020-10-25 11:42",
					EventDate: "2010-10-25",
					Typ:       "same",
				}
				eventDetailService.EXPECT().ValidateFormEvent(eventDetailForm).Return(errors.New("invalid link"))

				s.adminEvent(w, req)
				resp := w.Result()

				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})

			Convey("failure: post /admin-event-detail db error creating record invalid html", func() {
				s.views.AdminEventDetail = "testdata/invalid-admin-event-detail.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=http://www.google.com&createdAt=2020-10-25 11:42&eventDate=2010-10-25&eventType=same")
				req := httptest.NewRequest("POST", "/admin-event-detail", body)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				eventDetailForm := model.EventDetailForm{
					Title:     "title1",
					Info:      "info1",
					Link:      "http://www.google.com",
					CreatedAt: "2020-10-25 11:42",
					EventDate: "2010-10-25",
					Typ:       "same",
				}
				eventDetailService.EXPECT().ValidateFormEvent(eventDetailForm).Return(nil)
				eventDetailService.EXPECT().CreateNewEventDetail(eventDetailForm).Return(int64(0), errors.New("db error"))
				s.adminEvent(w, req)
				resp := w.Result()

				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})

			Convey("failure: post /admin-event-detail, empty body", func() {
				s.views.AdminEventDetail = "../../views/admin-event-detail.gtpl"
				req := httptest.NewRequest("POST", "/admin-event-detail", nil)
				req.Body = nil
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				infoForm := model.EventDetailForm{
					Title:     "title1",
					Info:      "info1",
					EventDate: "2020-10-25",
					Link:      "http://www.google.com",
					Typ:       "same",
					CreatedAt: "2020-10-27 08:08",
				}
				eventDetailService.EXPECT().ValidateFormEvent(infoForm).Return(nil)
				eventDetailService.EXPECT().CreateNewEventDetail(infoForm).Return(int64(10), nil)

				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				s.adminEvent(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})
		})

	})
}
