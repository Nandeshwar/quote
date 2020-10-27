package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"quote/pkg/service"
	"quote/pkg/service/quote"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdminPage(t *testing.T) {
	Convey("Test Admin Page", t, func() {
		w := httptest.NewRecorder()
		s := NewServer(0, ImageSize{}, "abc", 1, service.InfoEventService{}, quote.NewQuoteService())
		Convey("success: GET API", func() {
			req := httptest.NewRequest("GET", "/admin", nil)
			s.views.Admin = "../../views/admin.gtpl"
			session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

			session.Values["authenticated"] = true
			// session will be expired in given seconds
			session.Options.MaxAge = s.sessionExpireSeconds
			s.admin(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
		})

		Convey("failure: GET API - invalid session", func() {
			req := httptest.NewRequest("GET", "/admin", nil)
			s.views.Admin = "../../views/admin.gtpl"
			session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

			session.Values["authenticated"] = false
			// session will be expired in given seconds
			session.Options.MaxAge = s.sessionExpireSeconds
			s.admin(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusForbidden)
		})

		Convey("failure: GET API - wrong cookie name", func() {
			req := httptest.NewRequest("GET", "/admin", nil)
			s.views.Admin = "../../views/admin.gtpl"
			session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie1")

			session.Values["authenticated"] = true
			// session will be expired in given seconds
			session.Options.MaxAge = s.sessionExpireSeconds
			s.admin(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusForbidden)
		})

		Convey("failure: GET API - admin view not found", func() {
			req := httptest.NewRequest("GET", "/admin", nil)
			s.views.Admin = "../../views/admin1.gtpl"
			session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

			session.Values["authenticated"] = true
			// session will be expired in given seconds
			session.Options.MaxAge = s.sessionExpireSeconds
			s.admin(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusNotFound)
		})

		Convey("failure: GET API - cookie name should not be empty", func() {
			req := httptest.NewRequest("GET", "/admin", nil)
			s.views.Admin = "../../views/admin1.gtpl"
			s.cookieName = ""

			s.admin(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusForbidden)
		})

		Convey("failure: GET API - invalid html file", func() {
			req := httptest.NewRequest("GET", "/admin", nil)
			s.views.Admin = "testdata/invalid-html-file.gtpl"

			session, _ := s.sessionCookieStore.Get(req, "nandeshwar-quote-cookie")

			session.Values["authenticated"] = true
			// session will be expired in given seconds
			session.Options.MaxAge = s.sessionExpireSeconds

			s.admin(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusInternalServerError)
		})
	})
}
