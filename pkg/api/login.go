package api

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func (s Server) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		user := r.Form["username"]
		password := r.Form[("password")]

		err := s.loginService.Login(user[0], password[0])
		if err != nil {
			logrus.WithError(err).Error("invalid login id and password")
			http.Redirect(w, r, "login", http.StatusSeeOther)
			return
		}

		session, _ := s.sessionCookieStore.Get(r, "cookie-name")

		session.Values["authenticated"] = true
		// session will be expired in given seconds
		session.Options.MaxAge = s.sessionExpireSeconds
		session.Save(r, w)

		t, _ := template.ParseFiles("./views/admin.gtpl")
		t.Execute(w, nil)

	}
}
