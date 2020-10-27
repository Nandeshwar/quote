package api

import (
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

func (s Server) admin(w http.ResponseWriter, r *http.Request) {
	session, err := s.sessionCookieStore.Get(r, s.cookieName)

	if err != nil {
		logrus.Error("error getting session info", err)
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		logrus.Error("session expired")
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if r.Method == "GET" {
		t, err := template.ParseFiles(s.views.Admin)
		if err != nil {
			logrus.Error("error parsing view admin", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			logrus.Error("error executing view Admin view", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
}
