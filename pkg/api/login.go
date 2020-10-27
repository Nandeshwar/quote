package api

import (
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func (s Server) login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t, err := template.ParseFiles(s.views.Login)
		if err != nil {
			logrus.Error("error parsing file ./views/login.gtpl error=", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			logrus.Error("error executing login view", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case "POST":
		err := r.ParseForm()
		if err != nil {
			logrus.Error("error parsing form", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user := r.Form["username"]
		password := r.Form[("password")]

		err = s.loginService.Login(user[0], password[0])
		if err != nil {
			logrus.WithError(err).Error("invalid login id and password")
			http.Redirect(w, r, "login", http.StatusSeeOther)
			return
		}

		session, err := s.sessionCookieStore.Get(r, s.cookieName)
		if err != nil {
			logrus.WithError(err).Errorf("error getting cookie name")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		session.Values["authenticated"] = true
		// session will be expired in given seconds
		session.Options.MaxAge = s.sessionExpireSeconds
		session.Save(r, w)

		t, err := template.ParseFiles(s.views.Admin)
		if err != nil {
			logrus.Error("error executing login view", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			logrus.Error("error executing login view", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
