package api

import (
	"errors"
	"net/http"
	"quote/pkg/service/quote"
	"strings"
	"testing"

	"net/http/httptest"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	mock_service "quote/pkg/service/mock"

	"quote/pkg/service"
)

func TestLogin(t *testing.T) {
	Convey("Test login page", t, func() {
		ctrl := gomock.NewController(t)
		loginService := mock_service.NewMockILogin(ctrl)

		w := httptest.NewRecorder()
		s := NewServer(0, ImageSize{}, "abc", 1, service.InfoEventService{}, quote.NewQuoteService())
		s.loginService = loginService

		Convey("success: Get", func() {
			s.views.Login = "../../views/login.gtpl"
			req := httptest.NewRequest("GET", "/login", nil)
			s.login(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
			So(w.Body.String(), ShouldContainSubstring, "http://localhost:1922/quotes-devotional")
		})

		Convey("failure: wrong html file location", func() {
			s.views.Login = "login.gtpl"
			req := httptest.NewRequest("GET", "/login", nil)
			s.login(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusNotFound)
		})

		Convey("success: post", func() {
			reader := strings.NewReader("username=user&password=password")
			s.views.Admin = "../../views/admin.gtpl"
			req := httptest.NewRequest("POST", "/login", reader)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			loginService.EXPECT().Login("user", "password").Return(nil)
			s.login(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusOK)
			So(w.Body.String(), ShouldContainSubstring, `1. Add Info`)
			So(w.Body.String(), ShouldContainSubstring, `2. Add Events`)
		})

		Convey("failure: wrong credentials", func() {
			reader := strings.NewReader("username=user&password=password")
			s.views.Admin = "../../views/admin.gtpl"
			req := httptest.NewRequest("POST", "/login", reader)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			loginService.EXPECT().Login("user", "password").Return(errors.New("user credentials failed"))
			s.login(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusSeeOther)
		})

		Convey("failure: put", func() {
			reader := strings.NewReader("username=user&password=password")
			s.views.Admin = "../../views/admin.gtpl"
			req := httptest.NewRequest("PUT", "/login", reader)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			loginService.EXPECT().Login("user", "password").Return(nil)
			s.server.Handler.ServeHTTP(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusMethodNotAllowed)
		})

		Convey("failure wrong path for view", func() {
			s.views.Login = "./login.gtpl"
			req := httptest.NewRequest("GET", "/login", nil)
			s.login(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusNotFound)
		})

		Convey("failure: post - empty cookie name", func() {
			reader := strings.NewReader("username=user&password=password")
			s.cookieName = ""
			s.views.Admin = "../../views/admin.gtpl"
			req := httptest.NewRequest("POST", "/login", reader)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			loginService.EXPECT().Login("user", "password").Return(nil)
			s.login(w, req)
			resp := w.Result()
			So(resp.StatusCode, ShouldEqual, http.StatusForbidden)
		})
	})
}
