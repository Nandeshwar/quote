package api

import (
	"errors"
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

func TestAdminInfo(t *testing.T) {
	Convey("Test Admin Info API", t, func() {
		w := httptest.NewRecorder()
		s := NewServer(0, 0, ImageSize{}, "abc", 1, service.InfoEventService{}, quote.NewQuoteService())
		Convey("GET API", func() {
			Convey("success: Get /admin-info", func() {
				s.views.AdminInfo = "../../views/admin-info.gtpl"
				req := httptest.NewRequest("GET", "/admin-info", nil)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminInfo(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "Title")
				So(w.Body.String(), ShouldContainSubstring, "Info")
				So(w.Body.String(), ShouldContainSubstring, "Links Pipeline(|) Separated")
				So(w.Body.String(), ShouldContainSubstring, "Created Date")
			})

			Convey("failure: Get /admin-info view not found", func() {
				s.views.AdminInfo = "/login1.gtpl"
				req := httptest.NewRequest("GET", "/admin-info", nil)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminInfo(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusNotFound)
			})

			Convey("failure: Get /admin-info authentication false", func() {
				s.views.AdminInfo = "../../views/admin-info.gtpl"
				req := httptest.NewRequest("GET", "/admin-info", nil)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = false
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminInfo(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusForbidden)
			})

			Convey("failure: Get /admin-info wrong cookie store ", func() {
				s.views.AdminInfo = "../../views/admin-info.gtpl"
				req := httptest.NewRequest("GET", "/admin-info", nil)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie2")

				session.Values["authenticated"] = false
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminInfo(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusForbidden)
			})

			Convey("failure: Get /admin-info empty cookie name", func() {
				req := httptest.NewRequest("GET", "/admin-info", nil)

				s.cookieName = ""

				s.adminInfo(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusForbidden)
			})

			Convey("failure: Get /admin-info invalid html", func() {
				s.views.AdminInfo = "testdata/invalid-html-file.gtpl"
				req := httptest.NewRequest("GET", "/admin-info", nil)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminInfo(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})
		})
		Convey("Post API", func() {
			ctrl := gomock.NewController(t)
			infoService := mock_service.NewMockIInfo(ctrl)
			s.infoService = infoService
			Convey("success: post /admin-info", func() {
				s.views.AdminInfo = "../../views/admin-info.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=http://www.google.com&createdAt=2020-10-25 11:42")
				req := httptest.NewRequest("POST", "/admin-info", body)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				infoForm := model.InfoForm{
					Title:     "title1",
					Info:      "info1",
					Link:      "http://www.google.com",
					CreatedAt: "2020-10-25 11:42",
				}
				infoService.EXPECT().ValidateForm(infoForm).Return(nil)
				infoService.EXPECT().GetInfoLinkIDs(infoForm.Link).Return([]int64{10}, nil)
				infoService.EXPECT().CreateNewInfo(req.Context(), infoForm).Return(int64(10), nil)

				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				s.adminInfo(w, req)
				resp := w.Result()

				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "10")
			})

			Convey("failure: post /admin-info missing field createdAt", func() {
				s.views.AdminInfo = "../../views/admin-info.gtpl"
				body := strings.NewReader("title=title1&info1=info&link=http://www.google.com&createdAt=2020-10-25 11:42")
				req := httptest.NewRequest("POST", "/admin-info", body)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				s.adminInfo(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusNotFound)
			})

			Convey("failure: post /admin-info invalid form", func() {
				s.views.AdminInfo = "../../views/admin-info.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=abc&createdAt=2020-10-25 11:42")
				req := httptest.NewRequest("POST", "/admin-info", body)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				infoForm := model.InfoForm{
					Title:     "title1",
					Info:      "info1",
					Link:      "abc",
					CreatedAt: "2020-10-25 11:42",
				}
				infoService.EXPECT().ValidateForm(infoForm).Return(errors.New("invalid link"))

				s.adminInfo(w, req)
				//s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()

				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "validation error. check log")
			})

			Convey("failure: post /admin-info db error creating record", func() {
				s.views.AdminInfo = "../../views/admin-info.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=http://google.com/&createdAt=2020-10-25 11:42")
				req := httptest.NewRequest("POST", "/admin-info", body)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				infoForm := model.InfoForm{
					Title:     "title1",
					Info:      "info1",
					Link:      "http://google.com/",
					CreatedAt: "2020-10-25 11:42",
				}
				infoService.EXPECT().ValidateForm(infoForm).Return(nil)
				infoService.EXPECT().GetInfoLinkIDs(infoForm.Link).Return([]int64{10}, nil)
				infoService.EXPECT().CreateNewInfo(req.Context(), infoForm).Return(int64(0), errors.New("db error"))
				s.adminInfo(w, req)
				resp := w.Result()

				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(w.Body.String(), ShouldContainSubstring, "Did not create record. error. check log")
			})

			Convey("failure: post /admin-info invalid  html", func() {
				s.views.AdminInfo = "testdata/invalid-admin-info.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=http://www.google.com&createdAt=2020-10-25 11:42")
				req := httptest.NewRequest("POST", "/admin-info", body)
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				infoForm := model.InfoForm{
					Title:     "title1",
					Info:      "info1",
					Link:      "http://www.google.com",
					CreatedAt: "2020-10-25 11:42",
				}
				infoService.EXPECT().ValidateForm(infoForm).Return(nil)
				infoService.EXPECT().GetInfoLinkIDs(infoForm.Link).Return([]int64{10}, nil)
				infoService.EXPECT().CreateNewInfo(req.Context(), infoForm).Return(int64(10), nil)

				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				s.adminInfo(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})

			Convey("failure: post /admin-info db error creating record invalid html", func() {
				s.views.AdminInfo = "testdata/invalid-admin-info.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=http://google.com/&createdAt=2020-10-25 11:42")
				req := httptest.NewRequest("POST", "/admin-info", body)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				infoForm := model.InfoForm{
					Title:     "title1",
					Info:      "info1",
					Link:      "http://google.com/",
					CreatedAt: "2020-10-25 11:42",
				}
				infoService.EXPECT().ValidateForm(infoForm).Return(nil)
				infoService.EXPECT().GetInfoLinkIDs(infoForm.Link).Return([]int64{10}, nil)
				infoService.EXPECT().CreateNewInfo(req.Context(), infoForm).Return(int64(0), errors.New("db error"))
				s.adminInfo(w, req)
				resp := w.Result()

				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})

			Convey("failure: post /admin-info invalid form- invalid html", func() {
				s.views.AdminInfo = "testdata/invalid-admin-info.gtpl"
				body := strings.NewReader("title=title1&info=info1&link=abc&createdAt=2020-10-25 11:42")
				req := httptest.NewRequest("POST", "/admin-info", body)
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				infoForm := model.InfoForm{
					Title:     "title1",
					Info:      "info1",
					Link:      "abc",
					CreatedAt: "2020-10-25 11:42",
				}
				infoService.EXPECT().ValidateForm(infoForm).Return(errors.New("invalid link"))

				s.adminInfo(w, req)
				//s.server.Handler.ServeHTTP(w, req)
				resp := w.Result()

				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})

			Convey("failure: post /admin-info - empty body - parse error", func() {
				s.views.AdminInfo = "../../views/admin-info.gtpl"
				req := httptest.NewRequest("POST", "/admin-info", nil)
				req.Body = nil
				session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

				session.Values["authenticated"] = true
				// session will be expired in given seconds
				session.Options.MaxAge = s.sessionExpireSeconds

				infoForm := model.InfoForm{
					Title:     "title1",
					Info:      "info1",
					Link:      "http://www.google.com",
					CreatedAt: "2020-10-25 11:42",
				}
				infoService.EXPECT().ValidateForm(infoForm).Return(nil)
				infoService.EXPECT().CreateNewInfo(req.Context(), infoForm).Return(int64(10), nil)

				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				s.adminInfo(w, req)
				resp := w.Result()
				So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
			})
		})

	})
}
